import React, { useState } from 'react';
import { X, Save, Eye, EyeOff } from 'lucide-react';
import type { ProgramCredential } from '@/shared/types';
import type { CreateCredentialRequest, UpdateCredentialRequest } from '../types';

interface CredentialFormProps {
  credential?: ProgramCredential;  // undefined = create, populado = edit
  onSubmit: (data: CreateCredentialRequest | UpdateCredentialRequest) => void;
  onCancel: () => void;
  isLoading?: boolean;
}

// Form data type that includes all possible fields
interface FormData {
  name: string;
  user: string;
  password: string;
  access: string;
  notes: string;
  active: boolean;
}

const CredentialForm: React.FC<CredentialFormProps> = ({
  credential,
  onSubmit,
  onCancel,
  isLoading = false,
}) => {
  const [formData, setFormData] = useState<FormData>({
    name: credential?.name || '',
    user: credential?.user || '',
    password: credential?.password || '',
    access: credential?.access || '',
    notes: credential?.notes || '',
    active: credential?.active ?? true,
  });

  const [errors, setErrors] = useState<Record<string, string>>({});
  const [showPassword, setShowPassword] = useState(false);

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    // Name: obrigatório, mínimo 2 caracteres
    if (!formData.name || formData.name.trim().length < 2) {
      newErrors.name = 'Nome deve ter pelo menos 2 caracteres';
    }

    // User: opcional, se preenchido não pode ser vazio
    if (formData.user && formData.user.trim().length === 0) {
      newErrors.user = 'Usuário não pode ser vazio se preenchido';
    }

    // Access URL: opcional, valida URL pattern se preenchido
    if (formData.access && formData.access.trim().length > 0) {
        const urlPattern = /^(https?:\/\/)?([\da-z\.-]+)(:\d+)?(\/[\w\-._~:/?#[\]@!$&'()*+,;=]*)*$/i;
        if (!urlPattern.test(formData.access)) {
        newErrors.access = 'URL de acesso inválida';
      }
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();

    if (!validateForm()) {
      return;
    }

    // Build the appropriate request type based on whether we're editing
    if (credential) {
      const updateData: UpdateCredentialRequest = {
        name: formData.name,
        user: formData.user || undefined,
        password: formData.password || undefined,
        access: formData.access || undefined,
        notes: formData.notes || undefined,
        active: formData.active,
      };
      onSubmit(updateData);
    } else {
      const createData: CreateCredentialRequest = {
        name: formData.name,
        user: formData.user || undefined,
        password: formData.password || undefined,
        access: formData.access || undefined,
        notes: formData.notes || undefined,
        active: true, // Default true for create
      };
      onSubmit(createData);
    }
  };

  const handleChange = (field: keyof FormData, value: string | boolean) => {
    setFormData(prev => ({ ...prev, [field]: value }));
    // Clear error for this field when user starts typing
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: '' }));
    }
  };

  const isEditing = !!credential;

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4 animate-in fade-in duration-200">
      <div className="bg-[#1a1d24] rounded-2xl border border-white/10 shadow-2xl w-full max-w-lg max-h-[90vh] overflow-y-auto animate-in zoom-in-95 slide-in-from-bottom-4 duration-300">
        {/* Header */}
        <div className="flex items-center justify-between px-6 py-4 border-b border-gray-800">
          <h2 className="text-xl font-bold text-gray-100">
            {isEditing ? 'Editar Credencial' : 'Nova Credencial'}
          </h2>
          <button
            onClick={onCancel}
            className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
          >
            <X size={24} />
          </button>
        </div>

        {/* Form */}
        <form onSubmit={handleSubmit} className="p-6 space-y-5">
          {/* Name */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">
              Nome <span className="text-rose-400">*</span>
            </label>
            <input
              type="text"
              value={formData.name}
              onChange={(e) => handleChange('name', e.target.value)}
              className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                errors.name ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
              }`}
              placeholder="Ex: LinkedIn, Facebook, Google Ads"
            />
            {errors.name && <p className="text-xs text-rose-400">{errors.name}</p>}
          </div>

          {/* User */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">
              Usuário
            </label>
            <input
              type="text"
              value={formData.user}
              onChange={(e) => handleChange('user', e.target.value)}
              className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                errors.user ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
              }`}
              placeholder="Email ou nome de usuário"
            />
            {errors.user && <p className="text-xs text-rose-400">{errors.user}</p>}
          </div>

          {/* Password */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">
              Senha
            </label>
            <div className="relative">
              <input
                type={showPassword ? 'text' : 'password'}
                value={formData.password}
                onChange={(e) => handleChange('password', e.target.value)}
                className="w-full px-4 py-2.5 pr-12 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors"
                placeholder="Senha de acesso"
              />
              <button
                type="button"
                onClick={() => setShowPassword(!showPassword)}
                className="absolute right-3 top-1/2 -translate-y-1/2 text-gray-500 hover:text-gray-300 transition-colors"
              >
                {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
              </button>
            </div>
          </div>

          {/* Access URL */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">
              URL de Acesso
            </label>
            <input
              type="text"
              value={formData.access}
              onChange={(e) => handleChange('access', e.target.value)}
              className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                errors.access ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
              }`}
              placeholder="https://exemplo.com"
            />
            {errors.access && <p className="text-xs text-rose-400">{errors.access}</p>}
          </div>

          {/* Notes */}
          <div className="space-y-2">
            <label className="block text-sm font-medium text-gray-300">
              Observações
            </label>
            <textarea
              value={formData.notes}
              onChange={(e) => handleChange('notes', e.target.value)}
              rows={3}
              className="w-full px-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors resize-none"
              placeholder="Informações adicionais sobre a credencial..."
            />
          </div>

          {/* Active Status - Only show when editing */}
          {isEditing && (
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Status
              </label>
              <div className="flex items-center gap-3">
                <label className="flex items-center gap-2 cursor-pointer">
                  <input
                    type="radio"
                    name="active"
                    checked={formData.active === true}
                    onChange={() => handleChange('active', true)}
                    className="w-4 h-4 text-emerald-500 focus:ring-emerald-500 focus:ring-offset-gray-900"
                  />
                  <span className="text-sm text-gray-300">Ativo</span>
                </label>
                <label className="flex items-center gap-2 cursor-pointer">
                  <input
                    type="radio"
                    name="active"
                    checked={formData.active === false}
                    onChange={() => handleChange('active', false)}
                    className="w-4 h-4 text-rose-500 focus:ring-rose-500 focus:ring-offset-gray-900"
                  />
                  <span className="text-sm text-gray-300">Inativo</span>
                </label>
              </div>
            </div>
          )}

          {/* Footer */}
          <div className="flex items-center justify-end gap-3 pt-4 border-t border-gray-800">
            <button
              type="button"
              onClick={onCancel}
              disabled={isLoading}
              className="px-6 py-2.5 bg-gray-800 hover:bg-gray-700 text-white rounded-lg text-sm font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
            >
              Cancelar
            </button>
            <button
              type="submit"
              disabled={isLoading}
              className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg text-sm font-medium transition-colors disabled:opacity-50 disabled:cursor-not-allowed flex items-center gap-2 shadow-lg"
            >
              <Save size={16} />
              {isLoading ? 'Salvando...' : isEditing ? 'Atualizar' : 'Criar'}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default CredentialForm;
