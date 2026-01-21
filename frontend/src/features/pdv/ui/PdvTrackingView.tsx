import React, { useState } from 'react';
import type { PdvPost, RecurrentPdv } from '@/shared/types';
import {
  Plus,
  Filter,
  CalendarDays,
  Instagram,
  Facebook,
  Linkedin,
  Link as LinkIcon,
  CheckCircle2,
  Clock,
  ExternalLink,
  List,
  AlertCircle,
  MapPin,
  Store,
  Save
} from 'lucide-react';
import { usePdvPosts, useRecurrentPdvs, usePdvMutations, useRepresentatives } from '../hooks';
import type { PdvTab, PdvPlatform } from '../types';
import { useUIStore } from '@/shared/store/uiStore';

// WhatsApp Icon Component
const WhatsAppIcon = ({ size = 14, className = "" }: { size?: number, className?: string }) => (
    <svg
        xmlns="http://www.w3.org/2000/svg"
        width={size}
        height={size}
        viewBox="0 0 24 24"
        fill="currentColor"
        className={className}
    >
        <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.506.709.312 1.262.497 1.696.635.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413Z"/>
    </svg>
);

const PdvTrackingView: React.FC = () => {
  const [activeTab, setActiveTab] = useState<PdvTab>('posts');
  const { data: postsData = { posts: [], meta: { total: 0, page: 1, limit: 10, total_pages: 1 } } } = usePdvPosts();
  const { data: pdvListData = { pdvs: [], meta: { total: 0, page: 1, limit: 10, total_pages: 1 } } } = useRecurrentPdvs();
  const { data: representatives = [], isLoading: isLoadingReps } = useRepresentatives();
  const [isAdding, setIsAdding] = useState(false);
  const [isAddingPdv, setIsAddingPdv] = useState(false);

  // Filter States
  const { selectedMonth } = useUIStore();
  const [selectedRep, setSelectedRep] = useState('Todos');
  const [missingMonth, setMissingMonth] = useState('NOV'); // Default for missing tab
  
  // Post Form States
  const [newPostRepUUID, setNewPostRepUUID] = useState('');
  const [newPostPdv, setNewPostPdv] = useState('');
  const [newPostDate, setNewPostDate] = useState('');
  const [newPostPlatform, setNewPostPlatform] = useState<PdvPlatform>('instagram');
  const [newPostLink, setNewPostLink] = useState('');
  
  // PDV Form States
  const [newPdvName, setNewPdvName] = useState('');
  const [newPdvRepUUID, setNewPdvRepUUID] = useState('');
  const [newPdvCity, setNewPdvCity] = useState('');
  
  const { createPost, createRecurrent } = usePdvMutations();

  // Helper function to get representative name from UUID
  const getRepresentativeName = (uuid: string) => {
    const rep = representatives.find(r => r.uuid === uuid);
    return rep?.name || uuid;
  };

  const handleAddSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    const newPost = {
      representative_uuid: newPostRepUUID,
      pdv_name: newPostPdv,
      post_date: newPostDate,
      platform: newPostPlatform,
      link: newPostLink || undefined,
    };
    createPost(newPost, {
      onSuccess: () => {
        setIsAdding(false);
        // Reset form
        setNewPostPdv('');
        setNewPostDate('');
        setNewPostLink('');
      }
    });
  };
  
  const handleAddPdvSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if(!newPdvName) return;
    
    const newPdv = {
      name: newPdvName,
      representative_uuid: newPdvRepUUID,
      city: newPdvCity || undefined,
    };
    
    createRecurrent(newPdv, {
      onSuccess: () => {
        setIsAddingPdv(false);
        setNewPdvName('');
        setNewPdvCity('');
      }
    });
  };
  
  const openDatePicker = (e: React.MouseEvent<HTMLInputElement>) => {
    try {
        if (e.currentTarget && 'showPicker' in e.currentTarget) {
           // @ts-ignore
           e.currentTarget.showPicker();
        }
    } catch (error) {}
  };
  
  const filteredPosts = postsData.posts.filter(post => {
      const matchRep = selectedRep === 'Todos' || post.representative_uuid === selectedRep;
      const matchMonth = selectedMonth === 'Todos' || post.month === selectedMonth;
      return matchRep && matchMonth;
  });
  
  const getPlatformIcon = (platform: string) => {
      switch(platform) {
          case 'instagram': return <Instagram size={16} className="text-pink-400" />;
          case 'facebook': return <Facebook size={16} className="text-blue-400" />;
          case 'linkedin': return <Linkedin size={16} className="text-blue-600" />;
          case 'youtube': return <LinkIcon size={16} className="text-red-500" />;
          case 'tiktok': return <LinkIcon size={16} className="text-gray-400" />;
          default: return <LinkIcon size={16} className="text-gray-400" />;
      }
  };
  
  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
          
        {/* Unified Control Bar */}
        <div className="flex flex-col md:flex-row justify-between items-center mb-6 gap-4">
              
            {/* Filters (Left Side) */}
            <div className="flex flex-wrap gap-4 w-full md:w-auto items-center">
                {activeTab === 'posts' ? (
                    <>
                        <div className="relative">
                            <select 
                                value={selectedRep}
                                onChange={(e) => setSelectedRep(e.target.value)}
                                className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                                disabled={isLoadingReps}
                            >
                                <option value="">Todos Representantes</option>
                                {representatives.map(rep => <option key={rep.uuid} value={rep.uuid}>{rep.name}</option>)}
                            </select>
                            <Filter size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>
                        
                        <div className="relative">
                            <select
                                value={selectedMonth}
                                className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            >
                                <option value="Todos">Todos Meses</option>
                                <option value="NOV">Novembro</option>
                                <option value="OUT">Outubro</option>
                            </select>
                            <CalendarDays size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>
                    </>
                ) : (
                    // Filters for Missing View
                    <div className="flex items-center gap-4">
                        <div className="relative">
                            <select 
                                value={missingMonth}
                                onChange={(e) => setMissingMonth(e.target.value)}
                                className="appearance-none bg-[#1a1d24] border border-gray-700 text-gray-300 rounded-lg pl-10 pr-8 py-2 text-sm focus:outline-none focus:ring-1 focus:ring-[#1e5144]"
                            >
                                <option value="NOV">Novembro</option>
                                <option value="OUT">Outubro</option>
                            </select>
                            <CalendarDays size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                        </div>
                        <span className="text-gray-500 text-xs md:text-sm">PDVs sem postagem no mês selecionado</span>
                    </div>
                )}
            </div>
  
            {/* Actions (Right Side) */}
            <div className="flex items-center gap-3 w-full md:w-auto justify-end">
                {/* Tabs Switcher moved here */}
                <div className="flex gap-2 bg-[#1a1d24] p-1 rounded-lg border border-gray-700">
                    <button 
                        onClick={() => setActiveTab('posts')}
                        className={`px-3 py-1.5 rounded-md text-sm font-medium transition-colors flex items-center gap-2 ${activeTab === 'posts' ? 'bg-[#1e5144] text-white shadow-sm' : 'text-gray-400 hover:text-white'}`}
                    >
                        <List size={16} /> Histórico
                    </button>
                    <button 
                        onClick={() => setActiveTab('missing')}
                        className={`px-3 py-1.5 rounded-md text-sm font-medium transition-colors flex items-center gap-2 ${activeTab === 'missing' ? 'bg-[#1e5144] text-white shadow-sm' : 'text-gray-400 hover:text-white'}`}
                    >
                        <AlertCircle size={16} /> Pendentes
                    </button>
                </div>
  
                {activeTab === 'posts' && (
                    <button 
                        onClick={() => setIsAdding(!isAdding)}
                        className="bg-[#1e5144] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#163c32] transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/30"
                    >
                        <Plus size={18} /> Novo Registro
                    </button>
                )}
  
                {activeTab === 'missing' && (
                    <button 
                        onClick={() => setIsAddingPdv(!isAddingPdv)}
                        className="bg-[#1e5144] text-white px-4 py-2 rounded-lg text-sm font-medium hover:bg-[#163c32] transition-colors flex items-center gap-2 shadow-lg shadow-[#1e5144]/30"
                    >
                        <Plus size={18} /> Novo PDV
                    </button>
                )}
            </div>
        </div> 
  
        {activeTab === 'posts' ? (
            <>
                {/* Add Form Panel */}
                {isAdding && (
                    <div className="bg-[#1a1d24] rounded-xl border border-gray-700 shadow-xl mb-6 animate-in slide-in-from-top-4 flex flex-col">
                        <div className="p-4 border-b border-gray-800 bg-[#20232b]">
                            <h3 className="font-bold text-gray-100">Novo Registro de Postagem</h3>
                        </div>
                        
                        <form onSubmit={handleAddSubmit} className="p-6 space-y-4">
                            <div>
                                <label className="block text-sm font-medium text-gray-400 mb-1">Representante</label>
                                <select
                                    value={newPostRepUUID}
                                    onChange={(e) => setNewPostRepUUID(e.target.value)}
                                    className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                    disabled={isLoadingReps}
                                >
                                    <option value="">Selecione o Representante...</option>
                                    {representatives.map(rep => <option key={rep.uuid} value={rep.uuid}>{rep.name}</option>)}
                                </select>
                            </div>
  
                            <div>
                                <label className="block text-sm font-medium text-gray-400 mb-1">
                                    Nome do PDV <span className="text-rose-500">*</span>
                                </label>
                                <input 
                                    type="text" 
                                    value={newPostPdv}
                                    onChange={(e) => setNewPostPdv(e.target.value)}
                                    placeholder="Ex: Elétrica Silva"
                                    className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] placeholder-gray-600"
                                    required
                                />
                            </div>
  
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Data da Postagem</label>
                                    <div className="date-input-wrapper relative">
                                        <input 
                                            type="date" 
                                            value={newPostDate}
                                            onChange={(e) => setNewPostDate(e.target.value)}
                                            onClick={openDatePicker}
                                            className="w-full pl-9 pr-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] cursor-pointer"
                                            required
                                        />
                                        <CalendarDays size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500 pointer-events-none" />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Plataforma</label>
                                    <select 
                                        value={newPostPlatform}
                                        onChange={(e) => setNewPostPlatform(e.target.value as any)}
                                        className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                    >
                                        <option value="instagram">Instagram</option>
                                        <option value="facebook">Facebook</option>
                                        <option value="linkedin">LinkedIn</option>
                                        <option value="youtube">YouTube</option>
                                        <option value="tiktok">TikTok</option>
                                    </select>
                                </div>
                            </div>
  
                            <div>
                                <label className="block text-sm font-medium text-gray-400 mb-1">
                                    Link da Postagem <span className="text-gray-500 font-normal ml-1">(Opcional)</span>
                                </label>
                                <input 
                                    type="url" 
                                    value={newPostLink}
                                    onChange={(e) => setNewPostLink(e.target.value)}
                                    placeholder="https://..."
                                    className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] placeholder-gray-600"
                                />
                            </div>
  
                            <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                                <button 
                                    type="button" 
                                    onClick={() => setIsAdding(false)}
                                    className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                                >
                                    Cancelar
                                </button>
                                <button 
                                    type="submit" 
                                    className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg"
                                >
                                    Registrar
                                </button>
                            </div>
                        </form>
                    </div>
                )}
  
                {/* Content List */}
                <div className="bg-[#1a1d24] rounded-xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col">
                    <div className="overflow-x-auto custom-scrollbar flex-grow">
                        <table className="w-full text-sm text-left">
                            <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
                                <tr>
                                    <th className="px-6 py-4">PDV</th>
                                    <th className="px-6 py-4">Representante</th>
                                    <th className="px-6 py-4">Cidade</th>
                                    <th className="px-6 py-4">Data</th>
                                    <th className="px-6 py-4">Plataforma</th>
                                    <th className="px-6 py-4 text-center">Status</th>
                                    <th className="px-6 py-4 text-center">Link</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-800">
                                {filteredPosts.length === 0 ? (
                                    <tr>
                                        <td colSpan={7} className="px-6 py-8 text-center text-gray-500">
                                                    Nenhum registro encontrado com os filtros atuais.
                                        </td>
                                    </tr>
                                ) : (
                                    filteredPosts.map(post => {
                                        // Lookup City info if available in Recurrent list, otherwise empty
                                        const city = pdvListData.pdvs.find(p => p.name === post.pdv_name)?.city || '-';
                                         
                                        return (
                                            <tr key={post.id} className="hover:bg-[#20232b] transition-colors">
                                                <td className="px-6 py-4 font-medium text-gray-200">
                                                    {post.pdv_name}
                                                </td>
                                                <td className="px-6 py-4 text-gray-400">
                                                    {getRepresentativeName(post.representative_uuid)}
                                                </td>
                                                <td className="px-6 py-4">
                                                    <span className="flex items-center gap-1.5 bg-[#0f1115] px-2 py-1 rounded w-fit text-xs border border-gray-800 text-gray-300">
                                                        <MapPin size={10} className="text-gray-500" />
                                                        {city}
                                                    </span>
                                                </td>
                                                <td className="px-6 py-4 text-gray-400">
                                                    {post.post_date.split('-').reverse().join('/')}
                                                </td>
                                                <td className="px-6 py-4">
                                                    <div className="flex items-center gap-2 text-gray-300">
                                                        {getPlatformIcon(post.platform)}
                                                        <span className="capitalize">{post.platform}</span>
                                                    </div>
                                                </td>
                                                <td className="px-6 py-4 text-center">
                                                    {post.status === 'approved' ? (
                                                        <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 text-xs font-medium">
                                                                    <CheckCircle2 size={12} /> Aprovado
                                                        </span>
                                                    ) : post.status === 'rejected' ? (
                                                        <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-red-500/10 text-red-400 border border-red-500/20 text-xs font-medium">
                                                                    <Clock size={12} /> Rejeitado
                                                        </span>
                                                    ) : post.status === 'published' ? (
                                                        <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-blue-500/10 text-blue-400 border border-blue-500/20 text-xs font-medium">
                                                                    <CheckCircle2 size={12} /> Publicado
                                                        </span>
                                                    ) : post.status === 'cancelled' ? (
                                                        <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-gray-500/10 text-gray-400 border border-gray-500/20 text-xs font-medium">
                                                                    <Clock size={12} /> Cancelado
                                                        </span>
                                                    ) : (
                                                        <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded-full bg-amber-500/10 text-amber-400 border border-amber-500/20 text-xs font-medium">
                                                                    <Clock size={12} /> Pendente
                                                        </span>
                                                    )}
                                                </td>
                                                <td className="px-6 py-4 text-center">
                                                    {post.link ? (
                                                        <a 
                                                            href={post.link} 
                                                            target="_blank" 
                                                            rel="noopener noreferrer"
                                                            className="text-[#1e5144] hover:text-emerald-400 transition-colors"
                                                        >
                                                            <ExternalLink size={16} />
                                                        </a>
                                                    ) : (
                                                        <span className="text-gray-600">-</span>
                                                    )}
                                                </td>
                                            </tr>
                                        );
                                    })
                                )}
                            </tbody>
                        </table>
                    </div>
                </div>
            </>
        ) : (
            <>
                {/* Add PDV Form Panel */}
                {isAddingPdv && (
                    <div className="bg-[#1a1d24] rounded-xl border border-gray-700 shadow-xl mb-6 animate-in slide-in-from-top-4 flex flex-col">
                        <div className="p-4 border-b border-gray-800 bg-[#20232b]">
                            <h3 className="font-bold text-gray-100 flex items-center gap-2">
                                <Store size={18} className="text-[#1e5144]" />
                                Novo PDV Parceiro
                            </h3>
                        </div>
                      
                        <form onSubmit={handleAddPdvSubmit} className="p-6 space-y-4">
                            <div>
                                <label className="block text-sm font-medium text-gray-400 mb-1">Nome do PDV <span className="text-rose-500">*</span></label>
                                <input 
                                    type="text" 
                                    value={newPdvName}
                                    onChange={(e) => setNewPdvName(e.target.value)}
                                    className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                    placeholder="Nome da Loja"
                                    required
                                />
                            </div>
  
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Representante</label>
                                    <select
                                        value={newPdvRepUUID}
                                        onChange={(e) => setNewPdvRepUUID(e.target.value)}
                                        className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                        disabled={isLoadingReps}
                                    >
                                        <option value="">Selecione o Representante...</option>
                                        {representatives.map(rep => <option key={rep.uuid} value={rep.uuid}>{rep.name}</option>)}
                                    </select>
                                </div>
                                <div>
                                    <label className="block text-sm font-medium text-gray-400 mb-1">Cidade</label>
                                    <input 
                                        type="text" 
                                        value={newPdvCity}
                                        onChange={(e) => setNewPdvCity(e.target.value)}
                                        className="w-full px-3 py-2 bg-[#0f1115] border border-gray-700 rounded-lg text-sm text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
                                        placeholder="Cidade - UF"
                                    />
                                </div>
                            </div>
  
                            <div className="pt-4 border-t border-gray-800 flex justify-end gap-3">
                                <button 
                                    type="button" 
                                    onClick={() => setIsAddingPdv(false)}
                                    className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium transition-colors"
                                >
                                    Cancelar
                                </button>
                                <button 
                                    type="submit" 
                                    className="px-4 py-2 bg-[#1e5144] text-white hover:bg-[#163c32] rounded-lg text-sm font-medium transition-colors shadow-lg flex items-center gap-2"
                                >
                                    <Save size={16} /> Salvar PDV
                                </button>
                            </div>
                        </form>
                    </div>
                )}
  
                <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-lg overflow-hidden flex-grow flex flex-col animate-in fade-in">
                    {/* Header removed here as filters are now in main toolbar */}
                  
                    <div className="overflow-x-auto custom-scrollbar flex-grow">
                        <table className="w-full text-sm text-left">
                            <thead className="bg-[#15171c] text-xs text-gray-400 uppercase font-semibold border-b border-gray-800 sticky top-0">
                                <tr>
                                    <th className="px-6 py-4">PDV</th>
                                    <th className="px-6 py-4">Representante</th>
                                    <th className="px-6 py-4">Cidade</th>
                                    <th className="px-6 py-4 text-center">Ação</th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-800">
                                {pdvListData.pdvs.filter(pdv => !postsData.posts.some(p =>
                                    p.pdv_name.toLowerCase().includes(pdv.name.toLowerCase()) &&
                                    p.month === missingMonth
                                )).map((pdv) => (
                                    <tr key={pdv.id} className="hover:bg-[#20232b] transition-colors group">
                                        <td className="px-6 py-4 text-gray-200 font-medium">{pdv.name}</td>
                                        <td className="px-6 py-4 text-gray-300">{getRepresentativeName(pdv.representative_uuid)}</td>
                                        <td className="px-6 py-4">
                                            <span className="flex items-center gap-1.5 bg-[#0f1115] px-2 py-1 rounded w-fit text-xs border border-gray-800 text-gray-300">
                                                <MapPin size={10} className="text-gray-500" />
                                                {pdv.city || '-'}
                                            </span>
                                        </td>
                                        <td className="px-6 py-4 text-center">
                                            <button className="inline-flex items-center gap-1.5 px-3 py-1.5 bg-green-500/10 text-green-500 rounded-full text-xs font-medium border border-green-500/20 hover:bg-green-500/20 transition-colors">
                                                <WhatsAppIcon size={14} /> Cobrar Postagem
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                </div>
            </>
        )}
    </div>
  );
};
 
export default PdvTrackingView;
