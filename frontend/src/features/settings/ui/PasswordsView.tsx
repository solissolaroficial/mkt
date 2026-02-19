import React, { useState } from 'react';
import { Copy, Check, ExternalLink, Key, Lock, Globe, Plus, Edit2, Trash2 } from 'lucide-react';
import { useCredentials, useCredentialMutations } from '../hooks/useSettings';
import type { ProgramCredential } from '@/shared/types';
import type { CreateCredentialRequest, UpdateCredentialRequest } from '../types';
import CredentialForm from './CredentialForm';
import ConfirmationModal from '@/shared/components/ConfirmationModal';

const PasswordsView: React.FC = () => {
  const [copiedId, setCopiedId] = useState<string | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [selectedCredential, setSelectedCredential] = useState<ProgramCredential | undefined>(undefined);
  const [showDeleteConfirm, setShowDeleteConfirm] = useState<string | null>(null);
  
  const { data: credentials = [], isLoading } = useCredentials();
  const { createItem, updateItem, deleteItem, isCreating, isUpdating, isDeleting } = useCredentialMutations();

  const copyToClipboard = (text: string, id: string) => {
    navigator.clipboard.writeText(text);
    setCopiedId(id);
    setTimeout(() => setCopiedId(null), 2000);
  };

  // ==================== HANDLERS ====================

  const handleCreate = () => {
    setSelectedCredential(undefined);
    setShowForm(true);
  };

  const handleEdit = (credential: ProgramCredential) => {
    setSelectedCredential(credential);
    setShowForm(true);
  };

  const handleDelete = (id: string) => {
    setShowDeleteConfirm(id);
  };

  const handleConfirmDelete = () => {
    if (showDeleteConfirm) {
      deleteItem(showDeleteConfirm);
      setShowDeleteConfirm(null);
    }
  };

  const handleFormSubmit = (data: CreateCredentialRequest | UpdateCredentialRequest) => {
    if (selectedCredential) {
      // Edit mode
      updateItem(
        { id: selectedCredential.id, data: data as UpdateCredentialRequest },
        {
          onSuccess: () => {
            setShowForm(false);
            setSelectedCredential(undefined);
          },
        }
      );
    } else {
      // Create mode
      createItem(data as CreateCredentialRequest, {
        onSuccess: () => {
          setShowForm(false);
          setSelectedCredential(undefined);
        },
      });
    }
  };

  const handleFormCancel = () => {
    setShowForm(false);
    setSelectedCredential(undefined);
  };

  // ==================== RENDER ====================

  if (isLoading) {
    return (
      <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex items-center justify-center">
        <div className="text-gray-400">Carregando credenciais...</div>
      </div>
    );
  }

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="mb-6 flex items-center justify-between">
        <div>
          <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-3">
            <Lock className="text-emerald-500" />
            Acessos
          </h1>
          <p className="text-gray-500">Credenciais de acesso e ferramentas do Marketing</p>
        </div>
        <button
          onClick={handleCreate}
          className="flex items-center gap-2 px-4 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium transition-colors shadow-lg"
        >
          <Plus size={18} />
          Nova Credencial
        </button>
      </div>

      {/* Empty State */}
      {credentials.length === 0 ? (
        <div className="flex-1 flex flex-col items-center justify-center">
          <div className="p-6 bg-gray-800/50 rounded-full mb-4">
            <Lock size={48} className="text-gray-600" />
          </div>
          <h3 className="text-lg font-medium text-gray-300 mb-2">Nenhuma credencial encontrada</h3>
          <p className="text-gray-500 text-sm mb-6">Comece adicionando sua primeira credencial de acesso</p>
          <button
            onClick={handleCreate}
            className="flex items-center gap-2 px-6 py-3 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium transition-colors shadow-lg"
          >
            <Plus size={18} />
            Adicionar Primeira Credencial
          </button>
        </div>
      ) : (
        /* Cards Grid */
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 overflow-y-auto custom-scrollbar pb-4">
          {credentials.map((cred) => (
            <div
              key={cred.id}
              className={`bg-[#1a1d24] border border-gray-700/50 rounded-xl p-5 hover:border-emerald-500/30 transition-all group flex flex-col justify-between relative ${
                cred.active === false ? 'opacity-60' : ''
              }`}
            >
              {/* Action Buttons - Top Right */}
              <div className="absolute top-3 right-3 flex gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
                <button
                  onClick={() => handleEdit(cred)}
                  className="p-1.5 bg-gray-800 hover:bg-gray-700 text-gray-400 hover:text-white rounded border border-gray-700 transition-colors"
                  title="Editar"
                >
                  <Edit2 size={14} />
                </button>
                <button
                  onClick={() => handleDelete(cred.id)}
                  className="p-1.5 bg-gray-800 hover:bg-rose-900/50 text-gray-400 hover:text-rose-400 rounded border border-gray-700 hover:border-rose-700/50 transition-colors"
                  title="Excluir"
                >
                  <Trash2 size={14} />
                </button>
              </div>

              <div>
                <div className="flex justify-between items-start mb-4">
                  <div className="p-2 bg-gray-800 rounded-lg text-gray-300 border border-gray-700">
                    {cred.name === 'LinkedIn' || cred.name === 'Facebook' || cred.name === 'Instagram' || cred.name === 'TikTok' ? (
                      <Globe size={20} />
                    ) : (
                      <Key size={20} />
                    )}
                  </div>
                  {cred.access && !cred.access.includes(' ') && (
                    <a
                      href={cred.access.startsWith('http') ? cred.access : `https://${cred.access}`}
                      target="_blank"
                      rel="noopener noreferrer"
                      className="text-gray-500 hover:text-emerald-400 transition-colors"
                      title="Abrir Link"
                    >
                      <ExternalLink size={18} />
                    </a>
                  )}
                </div>

                <h3 className="font-bold text-lg text-white mb-3 pr-16">{cred.name}</h3>

                <div className="space-y-3">
                  {cred.user && (
                    <div>
                      <p className="text-[10px] text-gray-500 uppercase font-bold tracking-wider mb-0.5 flex justify-between">
                        Usuário
                      </p>
                      <div className="flex gap-2">
                        <div className="flex-grow text-sm text-gray-300 font-mono bg-black/20 p-1.5 rounded border border-gray-800 break-all">
                          {cred.user}
                        </div>
                        <button
                          onClick={() => copyToClipboard(cred.user!, `user-${cred.id}`)}
                          className="p-1.5 bg-gray-800 hover:bg-gray-700 text-gray-400 hover:text-white rounded border border-gray-700 transition-colors flex-shrink-0"
                          title="Copiar Usuário"
                        >
                          {copiedId === `user-${cred.id}` ? <Check size={14} className="text-emerald-500" /> : <Copy size={14} />}
                        </button>
                      </div>
                    </div>
                  )}

                  {cred.password ? (
                    <div>
                      <p className="text-[10px] text-gray-500 uppercase font-bold tracking-wider mb-0.5 flex justify-between">
                        Senha
                      </p>
                      <div className="flex gap-2">
                        <div className="flex-grow text-sm text-emerald-400 font-mono bg-emerald-900/10 p-1.5 rounded border border-emerald-900/30 truncate">
                          {cred.password}
                        </div>
                        <button
                          onClick={() => copyToClipboard(cred.password!, `pass-${cred.id}`)}
                          className="p-1.5 bg-gray-800 hover:bg-gray-700 text-gray-400 hover:text-white rounded border border-gray-700 transition-colors flex-shrink-0"
                          title="Copiar Senha"
                        >
                          {copiedId === `pass-${cred.id}` ? <Check size={14} className="text-emerald-500" /> : <Copy size={14} />}
                        </button>
                      </div>
                    </div>
                  ) : (
                    <div className="h-12 flex items-center">
                      <span className="text-xs text-gray-600 italic">Sem senha cadastrada</span>
                    </div>
                  )}

                  {cred.notes && (
                    <div className="mt-2 text-xs text-amber-400/80 bg-amber-900/10 p-2 rounded border border-amber-900/20">
                      {cred.notes}
                    </div>
                  )}
                </div>
              </div>

              {cred.access && (
                <div className="mt-4 pt-3 border-t border-gray-800">
                  <p className="text-[10px] text-gray-500 uppercase font-bold tracking-wider mb-0.5">Acesso</p>
                  <p className="text-xs text-gray-400 truncate" title={cred.access}>
                    {cred.access}
                  </p>
                </div>
              )}

              {/* Inactive Badge */}
              {cred.active === false && (
                <div className="absolute bottom-3 left-3">
                  <span className="text-[10px] text-rose-400 bg-rose-900/20 px-2 py-0.5 rounded border border-rose-900/30">
                    Inativo
                  </span>
                </div>
              )}
            </div>
          ))}
        </div>
      )}

      {/* Credential Form Modal */}
      {showForm && (
        <CredentialForm
          credential={selectedCredential}
          onSubmit={handleFormSubmit}
          onCancel={handleFormCancel}
          isLoading={isCreating || isUpdating}
        />
      )}

      {/* Delete Confirmation Modal */}
      <ConfirmationModal
        isOpen={showDeleteConfirm !== null}
        onClose={() => setShowDeleteConfirm(null)}
        onConfirm={handleConfirmDelete}
        title="Excluir Credencial"
        message="Tem certeza que deseja excluir esta credencial? Esta ação não pode ser desfeita."
        isPending={isDeleting}
        type="danger"
      />
    </div>
  );
};

export default PasswordsView;
