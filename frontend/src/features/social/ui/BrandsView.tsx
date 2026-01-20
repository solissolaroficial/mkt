import React, { useState } from 'react';
import { Plus, Trash2, Tag, Calendar, Loader2, CheckCircle2 } from 'lucide-react';
import { useBrands } from '../hooks/useBrands';
import { useBrandMutations } from '../hooks/useBrandMutations';
import type { Brand } from '@/shared/types/legacy.types';

interface BrandsViewProps {
  onRefresh?: () => void;
}

const BrandsView: React.FC<BrandsViewProps> = ({ onRefresh }) => {
  const [isAdding, setIsAdding] = useState(false);
  const [newBrandName, setNewBrandName] = useState('');
  const [showSuccess, setShowSuccess] = useState(false);

  const { data: brands = [], isLoading, error } = useBrands();
  const { createBrand, deleteBrand, isCreating, isDeleting } = useBrandMutations();

  const handleCreate = (e: React.FormEvent) => {
    e.preventDefault();
    if (!newBrandName.trim()) return;

    createBrand({ name: newBrandName.trim() }, {
      onSuccess: () => {
        setNewBrandName('');
        setIsAdding(false);
        setShowSuccess(true);
        setTimeout(() => setShowSuccess(false), 3000);
        if (onRefresh) onRefresh();
      },
      onError: (error) => {
        console.error('Erro ao criar marca:', error);
        alert('Erro ao criar marca. Tente novamente.');
      }
    });
  };

  const handleDelete = (id: string, name: string) => {
    if (window.confirm(`Tem certeza que deseja deletar a marca "${name}"?`)) {
      deleteBrand(id, {
        onSuccess: () => {
          if (onRefresh) onRefresh();
        },
        onError: (error) => {
          console.error('Erro ao deletar marca:', error);
          alert('Erro ao deletar marca. Tente novamente.');
        }
      });
    }
  };

  const formatDate = (dateString: string) => {
    try {
      const date = new Date(dateString);
      return date.toLocaleDateString('pt-BR', {
        day: '2-digit',
        month: '2-digit',
        year: 'numeric'
      });
    } catch {
      return dateString;
    }
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <Loader2 className="animate-spin text-gray-400" size={48} />
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 text-lg mb-2">Erro ao carregar marcas</p>
          <p className="text-gray-500 text-sm">Tente novamente mais tarde</p>
        </div>
      </div>
    );
  }

  return (
    <div className="h-full flex flex-col">
      {/* Header */}
      <div className="mb-6">
        <div className="flex justify-between items-center">
          <div>
            <h2 className="text-2xl font-bold text-gray-100 mb-1">Gerenciar Marcas</h2>
            <p className="text-gray-500 text-sm">
              {brands.length} {brands.length === 1 ? 'marca cadastrada' : 'marcas cadastradas'}
            </p>
          </div>
          {!isAdding && (
            <button
              onClick={() => setIsAdding(true)}
              className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white font-medium rounded-lg flex items-center gap-2 transition-colors shadow-lg shadow-[#1e5144]/20"
            >
              <Plus size={18} />
              Nova Marca
            </button>
          )}
        </div>
      </div>

      {/* Create Form */}
      {isAdding && (
        <div className="mb-6 p-4 bg-[#1a1d24] border border-[#1e5144]/30 rounded-xl">
          <form onSubmit={handleCreate} className="space-y-4">
            <div>
              <label className="block text-sm font-medium text-gray-400 mb-2">
                Nome da Marca <span className="text-rose-500">*</span>
              </label>
              <input
                type="text"
                value={newBrandName}
                onChange={(e) => setNewBrandName(e.target.value)}
                placeholder="Ex: Nike, Adidas, Puma..."
                className="w-full px-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                autoFocus
                required
              />
            </div>
            <div className="flex gap-3">
              <button
                type="submit"
                disabled={isCreating || !newBrandName.trim()}
                className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white font-medium rounded-lg flex items-center gap-2 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
              >
                {isCreating ? (
                  <>
                    <Loader2 size={16} className="animate-spin" />
                    Criando...
                  </>
                ) : (
                  <>
                    <Plus size={16} />
                    Criar Marca
                  </>
                )}
              </button>
              <button
                type="button"
                onClick={() => {
                  setIsAdding(false);
                  setNewBrandName('');
                }}
                className="px-4 py-2 bg-gray-700 hover:bg-gray-600 text-gray-300 font-medium rounded-lg transition-colors"
              >
                Cancelar
              </button>
            </div>
          </form>
        </div>
      )}

      {/* Success Message */}
      {showSuccess && (
        <div className="mb-4 p-3 bg-emerald-500/10 border border-emerald-500/30 rounded-lg flex items-center gap-2 text-emerald-400 animate-in fade-in slide-in-from-top-2">
          <CheckCircle2 size={18} />
          <span className="text-sm font-medium">Marca criada com sucesso!</span>
        </div>
      )}

      {/* Brands List */}
      {brands.length === 0 ? (
        <div className="flex-grow flex items-center justify-center">
          <div className="text-center">
            <Tag size={48} className="mx-auto mb-4 text-gray-600" />
            <p className="text-gray-400 text-lg mb-2">Nenhuma marca cadastrada</p>
            <p className="text-gray-600 text-sm">Clique em "Nova Marca" para começar</p>
          </div>
        </div>
      ) : (
        <div className="flex-grow overflow-y-auto custom-scrollbar">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
            {brands.map((brand) => (
              <div
                key={brand.id}
                className="bg-[#1a1d24] border border-gray-800 rounded-xl p-4 hover:border-gray-700 transition-colors group"
              >
                <div className="flex items-start justify-between mb-3">
                  <div className="flex items-center gap-3">
                    <div className="w-10 h-10 rounded-lg bg-[#1e5144]/20 flex items-center justify-center">
                      <Tag size={20} className="text-[#1e5144]" />
                    </div>
                    <div>
                      <h3 className="font-semibold text-gray-100">{brand.name}</h3>
                      <div className="flex items-center gap-1 text-xs text-gray-500 mt-1">
                        <Calendar size={12} />
                        {formatDate(brand.created_at)}
                      </div>
                    </div>
                  </div>
                  <button
                    onClick={() => handleDelete(brand.id, brand.name)}
                    disabled={isDeleting}
                    className="p-2 text-gray-500 hover:text-rose-400 hover:bg-rose-500/10 rounded-lg transition-colors opacity-0 group-hover:opacity-100 disabled:opacity-50"
                    title="Deletar marca"
                  >
                    {isDeleting ? (
                      <Loader2 size={16} className="animate-spin" />
                    ) : (
                      <Trash2 size={16} />
                    )}
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
};

export default BrandsView;
