import React, { useState, useEffect, useRef } from 'react';
import type { GiftItem, GiftTransaction, Notification } from '@/shared/types';
import { Package, AlertTriangle, Box, X, ArrowUpRight, ArrowDownRight, Plus, Save } from 'lucide-react';
import { useGiftItems, useGiftTransactions, useCreateGiftItem, useCreateTransaction } from '../hooks/useGifts';
import { REP_NAMES } from '@/shared/utils/legacy.constants';
import type { GiftsViewProps, TransactionTab } from '../types';

const GiftsView: React.FC<GiftsViewProps> = ({ onAddNotification }) => {
  // Query hooks
  const { data: itemsData = [] } = useGiftItems();
  const { data: transactionsData = [] } = useGiftTransactions();
  const createGiftItemMutation = useCreateGiftItem();
  const createTransactionMutation = useCreateTransaction();

  // State for items (initialized from query data)
  const [items, setItems] = useState<GiftItem[]>(itemsData);
  const [selectedGift, setSelectedGift] = useState<GiftItem | null>(null);
  
  // Initialize transactions with mock data AND generated initial states for other items
  const [transactions, setTransactions] = useState<GiftTransaction[]>(() => {
    const existingItems = new Set(transactionsData.map(t => t.itemName));
    const initialData: GiftTransaction[] = [];
    
    itemsData.forEach((item, idx) => {
      if (!existingItems.has(item.name)) {
         // Generate initial transaction based on static stock to ensure dynamic calculation works for everyone
         const isPositive = item.stock >= 0;
         initialData.push({
            id: `init-${idx}`,
            itemName: item.name,
            date: new Date().toLocaleDateString('pt-BR'), // Baseline date
            time: '08:00',
            quantity: Math.abs(item.stock),
            unit: 'unid.',
            type: isPositive ? 'in' : 'out',
            price: item.price,
            representative: isPositive ? undefined : 'Ajuste de Estoque Inicial'
         });
      }
    });
    
    return [...transactionsData, ...initialData];
  });
  
  // Modal States
  const [activeTab, setActiveTab] = useState<TransactionTab>('out');
  
  // Transaction Form States
  const [qty, setQty] = useState('');
  const [price, setPrice] = useState('');
  const [rep, setRep] = useState('');
  
  // Add Item Modal State
  const [isAddingItem, setIsAddingItem] = useState(false);
  const [newItemName, setNewItemName] = useState('');
  const [newItemPrice, setNewItemPrice] = useState('');
  
  // Notification Refs
  const notifiedRef = useRef(new Set<string>());
  const isFirstRender = useRef(true);
  
  // Helper to parse DD/MM/YYYY HH:MM to Date object for sorting
  const parseDateTime = (dateStr: string, timeStr: string) => {
    const parts = dateStr.split('/');
    if (parts.length !== 3) return new Date(0); // Fallback
    
    // Treat '-' time as 00:00
    const time = (timeStr && timeStr !== '-') ? timeStr : '00:00';
    return new Date(`${parts[2]}-${parts[1]}-${parts[0]}T${time}:00`);
  };
  
  // Dynamic Calculation Function
  const getCalculatedGiftData = (baseItem: GiftItem) => {
    const itemTransactions = transactions.filter(t => t.itemName === baseItem.name);
 
    let totalStock = 0;
    let lastPrice = baseItem.price;
    let lastInDate: Date | null = null;
 
    // Calculate Stock and find last price
    itemTransactions.forEach(t => {
        if (t.type === 'in') {
            totalStock += t.quantity;
            
            // Logic for "Last Unit Price" based on latest input date
            const tDate = parseDateTime(t.date, t.time);
            
            // Update price if this transaction is newer than the last one found
            if (!lastInDate || tDate >= lastInDate) {
                lastInDate = tDate;
                // Only update price if it's defined
                if (t.price !== undefined) {
                    lastPrice = t.price;
                }
            }
        } else {
            totalStock -= t.quantity;
        }
    });
 
    return {
        ...baseItem,
        stock: totalStock,
        price: lastPrice
    };
  };
  
  // Effect to check for low stock and notify
  useEffect(() => {
    if (!onAddNotification) return;
 
    items.forEach(baseItem => {
        const item = getCalculatedGiftData(baseItem);
        
        // Check for low stock (< 20)
        if (item.stock < 20) {
            if (!notifiedRef.current.has(item.name)) {
                // Only notify if NOT first render to avoid spam on page load
                // "itens que entrarem em estoque baixo" implies a transition or event
                if (!isFirstRender.current) {
                    onAddNotification({
                        id: `stock-${item.name}-${Date.now()}`,
                        type: 'system',
                        title: 'Estoque Baixo',
                        message: `Estoque Baixo, ${item.name} tem apenas ${item.stock} unid.`,
                        timestamp: 'Agora',
                        read: false
                    });
                }
                notifiedRef.current.add(item.name);
            }
        } else {
            // If stock recovered, remove from tracked set so we can notify again if it drops
            if (notifiedRef.current.has(item.name)) {
                notifiedRef.current.delete(item.name);
            }
        }
    });
 
    // Mark first render as done
    if (isFirstRender.current) {
        isFirstRender.current = false;
    }
 
  }, [transactions, onAddNotification, items]);
  
  const handleOpenModal = (item: GiftItem) => {
    // Recalculate to ensure modal has latest state
    const calculated = getCalculatedGiftData(item);
    setSelectedGift(calculated);
    setActiveTab('out'); 
    setQty('');
    setPrice('');
    setRep('');
  };
  
  const handleAddItem = (e: React.FormEvent) => {
      e.preventDefault();
      if (!newItemName) return;
 
      const newItem: GiftItem = {
          name: newItemName,
          stock: 0,
          price: Number(newItemPrice) || 0
      };
 
      createGiftItemMutation.mutate(newItem, {
        onSuccess: (createdItem) => {
          setItems([...items, createdItem]);
          setIsAddingItem(false);
          setNewItemName('');
          setNewItemPrice('');
        }
      });
  };
 
  const handleAddTransaction = (e: React.FormEvent) => {
    e.preventDefault();
    if (!selectedGift) return;
 
    if (!qty) return;
    if (activeTab === 'in' && !price) return;
    if (activeTab === 'out' && !rep) return;
 
    const newTx: GiftTransaction = {
      id: Date.now().toString(),
      itemName: selectedGift.name,
      date: new Date().toLocaleDateString('pt-BR'),
      time: new Date().toLocaleTimeString('pt-BR', { hour: '2-digit', minute: '2-digit' }),
      quantity: Number(qty),
      unit: 'pç',
      type: activeTab,
      price: activeTab === 'in' ? Number(price) : undefined,
      representative: activeTab === 'out' ? rep : undefined
    };
 
    createTransactionMutation.mutate(newTx, {
      onSuccess: () => {
        const newTransactions = [newTx, ...transactions];
        setTransactions(newTransactions);
        
        // Update selectedGift immediately for UI feedback
        let newStock = selectedGift.stock;
        let newPrice = selectedGift.price;
        
        if (activeTab === 'in') {
            newStock += Number(qty);
            // Assuming new entry is latest, update price
            newPrice = Number(price);
        } else {
            newStock -= Number(qty);
        }
        
        setSelectedGift({
            ...selectedGift,
            stock: newStock,
            price: newPrice
        });
 
        // Reset form
        setQty('');
        setPrice('');
        setRep('');
      }
    });
  };
  
  const parseDateSort = (dateStr: string) => {
    const parts = dateStr.split('/');
    if (parts.length === 3) {
      return `${parts[2]}-${parts[1]}-${parts[0]}`;
    }
    return dateStr;
  };
  
  // Filter and sort transactions for table
  const currentData = transactions
    .filter(t => t.itemName === selectedGift?.name && t.type === activeTab)
    .sort((a, b) => {
        const dateA = parseDateSort(a.date);
        const dateB = parseDateSort(b.date);
        if (dateA < dateB) return 1;
        if (dateA > dateB) return -1;
        // Secondary sort by time
        if (a.time && b.time && a.time !== '-' && b.time !== '-') {
            return a.time < b.time ?1 : -1;
        }
        return 0;
    });
 
  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-100">Brindes</h1>
          <p className="text-gray-500">Gestão de estoque e valores de brindes</p>
        </div>
        <button 
            onClick={() => setIsAddingItem(true)}
            className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/20"
        >
            <Plus size={18} /> Novo Brinde
        </button>
      </div> 
      
      {/* Grid of Mini KPIs */}
      <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-4 pb-6 overflow-y-auto custom-scrollbar">
        {items.map((baseItem, index) => {
            // CALCULATE DATA DYNAMICALLY
            const item = getCalculatedGiftData(baseItem);
 
            const isNegative = item.stock < 0;
            const isZero = item.stock === 0;
            const isLowStock = !isNegative && !isZero && item.stock < 20;
            
            let statusColor = 'bg-[#1a1d24] border-gray-700';
            let iconColor = 'bg-emerald-500/10 text-emerald-400 border-emerald-500/20';
            let Icon = Box;
 
            if (isNegative || isLowStock) {
                statusColor = 'bg-rose-900/10 border-rose-900/30';
                iconColor = 'bg-rose-500/10 text-rose-400 border-rose-500/20';
                Icon = AlertTriangle;
            } else if (isZero) {
                statusColor = 'bg-gray-800/40 border-gray-700';
                iconColor = 'bg-gray-700/50 text-gray-500 border-gray-600';
                Icon = Package;
            } 
            
            return (
                <div 
                    key={index} 
                    onClick={() => handleOpenModal(item)}
                    className={`rounded-xl border shadow-lg p-4 flex flex-col justify-between transition-all hover:-translate-y-1 hover:shadow-xl cursor-pointer ${statusColor} ${(isNegative || isLowStock) ? 'border-rose-500/20' : 'border-white/5'}`}
                >
                    <div className="flex justify-between items-start mb-2">
                        <div className={`p-2 rounded-lg border ${iconColor}`}>
                            <Icon size={18} />
                        </div>
                        {item.price > 0 && (
                            <span className="text-xs font-mono text-gray-400 bg-black/20 px-2 py-1 rounded">
                                {item.price.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}
                            </span>
                        )}
                    </div>
                    
                    <div>
                        <h3 className="text-sm font-medium text-gray-300 line-clamp-2 min-h-[40px] mb-1">
                            {item.name}
                        </h3>
                        <div className="flex items-end gap-2">
                            <span className={`text-2xl font-bold tracking-tight ${(isNegative || isLowStock) ? 'text-rose-500' : isZero ? 'text-gray-500' : 'text-white'}`}>
                                {item.stock}
                            </span>
                            <span className="text-xs text-gray-500 mb-1">unid.</span>
                        </div>
                    </div>
                </div>
            );
        })}
      </div> 
      
      {/* --- GIFT DETAIL MODAL --- */}
      {selectedGift && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setSelectedGift(null)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-4xl max-h-[90vh] rounded-2xl shadow-2xl border border-gray-700 flex flex-col overflow-hidden animate-in zoom-in-95 duration-200"
                onClick={e => e.stopPropagation()}
            >
                {/* Modal Header */}
                <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-start">
                    <div>
                        <h2 className="text-2xl font-bold text-white mb-1">{selectedGift.name}</h2>
                        <div className="flex items-center gap-4">
                            <span className="text-sm text-gray-400">Estoque Atual:</span>
                            <span className={`text-xl font-bold ${selectedGift.stock < 20 ? 'text-rose-500' : 'text-emerald-400'}`}>
                                {selectedGift.stock} <span className="text-sm font-normal text-gray-500">unid.</span>
                            </span>
                            {selectedGift.price > 0 && (
                                <span className="text-sm text-gray-400 border-l border-gray-700 pl-4">
                                    Valor Unitário: <span className="text-white font-mono">{selectedGift.price.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' })}</span>
                                </span>
                            )}
                        </div>
                    </div>
                    <button 
                        onClick={() => setSelectedGift(null)}
                        className="p-2 text-gray-400 hover:text-white hover:bg-gray-800 rounded-full transition-colors"
                    >
                        <X size={24} />
                    </button>
                </div> 
                
                {/* Tabs */}
                <div className="flex border-b border-gray-800 bg-[#1a1d24]">
                    <button 
                        onClick={() => setActiveTab('out')}
                        className={`flex-1 py-4 text-sm font-medium border-b-2 transition-all flex items-center justify-center gap-2 ${
                            activeTab === 'out' 
                            ? 'border-rose-500 text-rose-400 bg-rose-500/5' 
                            : 'border-transparent text-gray-500 hover:text-gray-300 hover:bg-white/5'
                        }`}
                    >
                        <ArrowUpRight size={18} /> Histórico de Saídas
                    </button>
                    <button 
                        onClick={() => setActiveTab('in')}
                        className={`flex-1 py-4 text-sm font-medium border-b-2 transition-all flex items-center justify-center gap-2 ${
                            activeTab === 'in' 
                            ? 'border-emerald-500 text-emerald-400 bg-emerald-500/5' 
                            : 'border-transparent text-gray-500 hover:text-gray-300 hover:bg-white/5'
                        }`}
                    >
                        <ArrowDownRight size={18} /> Histórico de Entradas
                    </button>
                </div> 
                
                {/* Add Data Form */}
                <div className="p-4 bg-[#15171c] border-b border-gray-800">
                    <form onSubmit={handleAddTransaction} className="flex flex-col md:flex-row gap-4 items-end">
                        <div className="w-full md:w-32">
                            <label className="text-xs text-gray-500 mb-1 block font-medium uppercase">Quantidade</label>
                            <input 
                                type="number" 
                                value={qty}
                                onChange={e => setQty(e.target.value)}
                                placeholder="0"
                                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-blue-500"
                            />
                        </div>
                        
                        {activeTab === 'in' ? (
                            <div className="w-full md:w-40">
                                <label className="text-xs text-gray-500 mb-1 block font-medium uppercase">Valor Unit. (R$)</label>
                                <input 
                                    type="number" 
                                    step="0.01"
                                    value={price}
                                    onChange={e => setPrice(e.target.value)}
                                    placeholder="0,00"
                                    className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-blue-500"
                                />
                            </div>
                        ) : (
                            <div className="w-full flex-grow">
                                <label className="text-xs text-gray-500 mb-1 block font-medium uppercase">Representante / Destino</label>
                                <div className="relative">
                                    <input 
                                        list="rep-options"
                                        value={rep}
                                        onChange={e => setRep(e.target.value)}
                                        placeholder="Selecione ou digite..."
                                        className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-blue-500"
                                    />
                                    <datalist id="rep-options">
                                        {REP_NAMES.map(r => <option key={r} value={r} />)}
                                        <option value="Marketing" />
                                        <option value="Comercial" />
                                    </datalist>
                                </div>
                            </div>
                        )}
                        
                        <button 
                            type="submit"
                            disabled={!qty || (activeTab === 'in' && !price) || (activeTab === 'out' && !rep)}
                            className="w-full md:w-auto px-6 py-2 bg-blue-600 hover:bg-blue-700 text-white rounded-lg text-sm font-medium transition-colors flex items-center justify-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
                        >
                            <Plus size={16} /> Adicionar
                        </button>
                    </form>
                </div> 
                
                {/* Data Table */}
                <div className="flex-grow overflow-y-auto custom-scrollbar bg-[#0f1115]">
                    <table className="w-full text-sm text-left">
                        <thead className="text-xs text-gray-400 uppercase bg-[#1a1d24] sticky top-0 border-b border-gray-800">
                            <tr>
                                <th className="px-6 py-3 font-semibold">Data</th>
                                <th className="px-6 py-3 font-semibold">Horário</th>
                                <th className="px-6 py-3 font-semibold text-right">Qtd</th>
                                {activeTab === 'in' ? (
                                    <th className="px-6 py-3 font-semibold text-right">Valor Unit.</th>
                                ) : (
                                    <th className="px-6 py-3 font-semibold">Destino</th>
                                )}
                            </tr>
                        </thead>
                        <tbody className="divide-y divide-gray-800">
                            {currentData.length > 0 ? (
                                currentData.map((t) => (
                                    <tr key={t.id} className="hover:bg-[#1a1d24] transition-colors">
                                        <td className="px-6 py-4 text-gray-300 font-medium whitespace-nowrap">{t.date}</td>
                                        <td className="px-6 py-4 text-gray-500 font-mono text-xs">{t.time}</td>
                                        <td className={`px-6 py-4 text-right font-bold ${t.type === 'in' ? 'text-emerald-400' : 'text-rose-400'}`}>
                                            {t.type === 'in' ? '+' : '-'}{t.quantity} <span className="text-[10px] text-gray-600 font-normal">{t.unit}</span>
                                        </td>
                                        {activeTab === 'in' ? (
                                            <td className="px-6 py-4 text-right text-gray-400 font-mono text-xs">
                                                {t.price && t.price > 0 ? t.price.toLocaleString('pt-BR', { style: 'currency', currency: 'BRL' }) : '-'}
                                            </td>
                                        ) : (
                                            <td className="px-6 py-4 text-gray-300 text-xs">
                                                {t.representative || '-'}
                                            </td>
                                        )}
                                    </tr>
                                ))
                            ) : (
                                <tr>
                                    <td colSpan={4} className="p-12 text-center text-gray-500">
                                        <div className="flex flex-col items-center gap-2">
                                            <div className="w-10 h-10 rounded-full bg-gray-800 flex items-center justify-center text-gray-600">
                                                {activeTab === 'in' ? <ArrowDownRight size={20} /> : <ArrowUpRight size={20} />}
                                            </div>
                                            <p>Nenhum registro de {activeTab === 'in' ? 'entrada' : 'saída'} encontrado.</p>
                                        </div>
                                    </td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
                
                {/* Footer Count */}
                <div className="p-3 bg-[#1a1d24] border-t border-gray-800 text-xs text-center text-gray-500">
                    Mostrando {currentData.length} registros
                </div>
            </div>
        </div>
      )} 
      
      {/* --- ADD NEW ITEM MODAL --- */}
      {isAddingItem && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={() => setIsAddingItem(false)}>
            <div 
                className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden animate-in zoom-in-95"
                onClick={e => e.stopPropagation()}
            >
                <div className="p-4 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                    <h3 className="font-bold text-gray-100 flex items-center gap-2">
                        <Package size={18} className="text-[#1e5144]" />
                        Novo Brinde
                    </h3>
                    <button onClick={() => setIsAddingItem(false)} className="text-gray-500 hover:text-white">
                        <X size={20} />
                    </button>
                </div>
                <form onSubmit={handleAddItem} className="p-6 space-y-4">
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Nome do Item</label>
                        <input 
                            type="text" 
                            value={newItemName}
                            onChange={e => setNewItemName(e.target.value)}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            placeholder="Ex: Camiseta Promocional"
                            required
                        />
                    </div>
                    <div>
                        <label className="block text-xs font-semibold text-gray-400 uppercase mb-1">Preço Unitário Estimado (R$)</label>
                        <input 
                            type="number" 
                            step="0.01"
                            value={newItemPrice}
                            onChange={e => setNewItemPrice(e.target.value)}
                            className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-sm text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            placeholder="0,00"
                        />
                    </div>
                    <div className="pt-2 flex justify-end gap-3">
                        <button type="button" onClick={() => setIsAddingItem(false)} className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm">Cancelar</button>
                        <button type="submit" className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium">Salvar</button>
                    </div>
                </form>
            </div>
        </div>
      )}
    </div>
  );
};
 
export default GiftsView;
