import React, { useState, useEffect } from 'react';
import type { Flow } from '@/shared/types';
import { X, Loader2 } from 'lucide-react';

interface FlowModalProps {
  isOpen: boolean;
  onClose: () => void;
  onSave: (data: { name: string; description?: string; color?: string }) => Promise<void>;
  editingFlow?: Flow | null;
}

export const FlowModal: React.FC<FlowModalProps> = ({
  isOpen,
  onClose,
  onSave,
  editingFlow,
}) => {
  const [name, setName] = useState(editingFlow?.name || '');
  const [description, setDescription] = useState(editingFlow?.description || '');
  const [color, setColor] = useState(editingFlow?.color || '#3B82F6');
  const [isSaving, setIsSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);

  // Sincronizar estados quando editingFlow muda
  useEffect(() => {
    if (editingFlow) {
      console.log('FlowModal: editingFlow received:', editingFlow);
      setName(editingFlow.name || '');
      setDescription(editingFlow.description || '');
      setColor(editingFlow.color || '#3B82F6');
    } else {
      // Resetar para valores padrão quando é criação
      setName('');
      setDescription('');
      setColor('#3B82F6');
      console.log('FlowModal: Creating new flow');
    }
    // Resetar erro ao abrir modal
    setError(null);
  }, [editingFlow, isOpen]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setError(null);

    if (!name.trim()) {
      setError('O nome é obrigatório');
      return;
    }

    console.log('FlowModal: handleSubmit called with:', { name, description, color, editingFlow });
    setIsSaving(true);
    try {
      await onSave({ name, description, color });
      console.log('FlowModal: onSave completed successfully');
      handleClose();
    } catch (err) {
      console.error('FlowModal: Error in onSave:', err);
      setError('Erro ao salvar fluxo. Tente novamente.');
    } finally {
      setIsSaving(false);
    }
  };

  const handleClose = () => {
    if (!isSaving) {
      setName(editingFlow?.name || '');
      setDescription(editingFlow?.description || '');
      setColor(editingFlow?.color || '#3B82F6');
      onClose();
    }
  };

  const colorOptions = [
    '#3B82F6', // Blue
    '#10B981', // Green
    '#F59E0B', // Amber
    '#EF4444', // Red
    '#8B5CF6', // Purple
    '#EC4899', // Pink
    '#06B6D4', // Cyan
    '#84CC16', // Lime
  ];

  if (!isOpen) return null;

  return (
    <div className="fixed inset-0 bg-black/80 z-50 flex justify-center items-center p-4 backdrop-blur-sm">
      <div className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-800 flex flex-col max-h-[85vh]">
        <div className="p-6 border-b border-gray-800 flex items-center justify-between">
          <h2 className="text-xl font-bold text-white">
            {editingFlow ? 'Editar Fluxo' : 'Novo Fluxo'}
          </h2>
          <button
            onClick={handleClose}
            className="text-gray-400 hover:text-white p-2 hover:bg-white/10 rounded-full transition-colors"
          >
            <X size={20} />
          </button>
        </div>

        <form onSubmit={handleSubmit} className="flex-1 overflow-y-auto p-6 flex flex-col">
          <div className="space-y-4 flex-1">
            {error && (
              <div className="mb-4 p-3 bg-red-500/10 border border-red-500/20 text-red-400 rounded-lg text-sm">
                {error}
              </div>
            )}

            <div>
              <label htmlFor="flow-name" className="text-sm font-medium text-gray-300 mb-2 block">
                Nome *
              </label>
              <input
                id="flow-name"
                type="text"
                value={name}
                onChange={(e) => setName(e.target.value)}
                required
                placeholder="Nome do fluxo"
                disabled={isSaving}
                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-2.5 text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed"
              />
            </div>

            <div>
              <label htmlFor="flow-description" className="text-sm font-medium text-gray-300 mb-2 block">
                Descrição
              </label>
              <textarea
                id="flow-description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Descrição opcional do fluxo"
                rows={3}
                disabled={isSaving}
                className="w-full bg-[#0f1115] border border-gray-700 rounded-lg px-4 py-2.5 text-sm text-gray-200 placeholder-gray-600 focus:outline-none focus:border-gray-600 disabled:opacity-50 disabled:cursor-not-allowed resize-none"
              />
            </div>

            <div>
              <label className="text-sm font-medium text-gray-300 mb-2 block">Cor</label>
              <div className="flex gap-2">
                {colorOptions.map((c) => (
                  <button
                    key={c}
                    type="button"
                    onClick={() => setColor(c)}
                    className={`w-8 h-8 rounded-full transition-all ${
                      color === c ? 'ring-2 ring-offset-2 ring-gray-400 scale-110' : 'opacity-60 hover:opacity-100'
                    }`}
                    style={{ backgroundColor: c }}
                    disabled={isSaving}
                  />
                ))}
              </div>
            </div>
          </div>

          <div className="pt-6 border-t border-gray-800 flex gap-3">
            <button
              type="button"
              onClick={handleClose}
              disabled={isSaving}
              className="flex-1 px-4 py-2.5 bg-[#20232b] text-gray-300 text-sm font-medium rounded-lg hover:bg-[#2a2e37] disabled:opacity-50 transition-colors"
            >
              Cancelar
            </button>
            <button
              type="submit"
              disabled={isSaving}
              className="flex-1 px-4 py-2.5 bg-[#3B82F6] text-white text-sm font-medium rounded-lg hover:bg-[#2563EB] disabled:opacity-50 transition-colors flex items-center justify-center gap-2"
            >
              {isSaving && <Loader2 size={16} className="animate-spin" />}
              {editingFlow ? 'Salvar' : 'Criar'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};
