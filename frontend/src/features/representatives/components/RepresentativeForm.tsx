import React from 'react';
import { X, Save } from 'lucide-react';
import type { Representative, CreateRepresentativeRequest, UpdateRepresentativeRequest } from '../types';

interface RepresentativeFormProps {
  representative?: Representative;
  onSubmit: (data: CreateRepresentativeRequest | UpdateRepresentativeRequest) => void;
  onCancel: () => void;
  isLoading?: boolean;
}

// Form data type that includes all possible fields
interface FormData {
  code: number;
  name: string;
  email: string;
  phone: string;
  company: string;
  region: string;
  city: string;
  attendant: string;
  active?: boolean;
}

const RepresentativeForm: React.FC<RepresentativeFormProps> = ({
  representative,
  onSubmit,
  onCancel,
  isLoading = false,
}) => {
  const [formData, setFormData] = React.useState<FormData>({
    code: representative?.code || 0,
    name: representative?.name || '',
    email: representative?.email || '',
    phone: representative?.phone || '',
    company: representative?.company || '',
    region: representative?.region || '',
    city: representative?.city || '',
    attendant: representative?.attendant || '',
    active: representative?.active ?? true,
  });

  const [errors, setErrors] = React.useState<Record<string, string>>({});

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    // Only validate code when creating (not editing)
    if (!representative && (!formData.code || formData.code < 100 || formData.code > 999)) {
      newErrors.code = 'Código deve estar entre 100 e 999';
    }

    if (!formData.name || formData.name.trim().length < 3) {
      newErrors.name = 'Nome deve ter pelo menos 3 caracteres';
    }

    if (!formData.email || !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(formData.email)) {
      newErrors.email = 'Email inválido';
    }

    if (!formData.company || formData.company.trim().length < 2) {
      newErrors.company = 'Empresa deve ter pelo menos 2 caracteres';
    }

    if (!formData.region || formData.region.trim().length < 2) {
      newErrors.region = 'Região deve ter pelo menos 2 caracteres';
    }

    if (!formData.city || formData.city.trim().length < 2) {
      newErrors.city = 'Cidade deve ter pelo menos 2 caracteres';
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
    if (representative) {
      const updateData: UpdateRepresentativeRequest = {
        name: formData.name,
        email: formData.email,
        phone: formData.phone,
        company: formData.company,
        region: formData.region,
        city: formData.city,
        attendant: formData.attendant,
        active: formData.active,
      };
      onSubmit(updateData);
    } else {
      const createData: CreateRepresentativeRequest = {
        code: formData.code,
        name: formData.name,
        email: formData.email,
        phone: formData.phone,
        company: formData.company,
        region: formData.region,
        city: formData.city,
        attendant: formData.attendant,
      };
      onSubmit(createData);
    }
  };

  const handleChange = (field: keyof FormData, value: any) => {
    setFormData(prev => ({ ...prev, [field]: value }));
    // Clear error for this field when user starts typing
    if (errors[field]) {
      setErrors(prev => ({ ...prev, [field]: '' }));
    }
  };

  const isEditing = !!representative;

  return (
    <div className="fixed inset-0 bg-black/50 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div className="bg-[#1a1d24] rounded-2xl border border-white/10 shadow-2xl w-full max-w-2xl max-h-[90vh] overflow-y-auto">
        {/* Header */}
        <div className="flex items-center justify-between px-6 py-4 border-b border-gray-800">
          <h2 className="text-xl font-bold text-gray-100">
            {isEditing ? 'Editar Representante' : 'Novo Representante'}
          </h2>
          <button
            onClick={onCancel}
            className="p-2 hover:bg-gray-800 rounded-full transition-colors text-gray-400 hover:text-white"
          >
            <X size={24} />
          </button>
        </div>

        {/* Form */}
        <form onSubmit={handleSubmit} className="p-6 space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            {/* Code - Only show when creating */}
            {!isEditing && (
              <div className="space-y-2">
                <label className="block text-sm font-medium text-gray-300">
                  Código <span className="text-rose-400">*</span>
                </label>
                <input
                  type="number"
                  min={100}
                  max={999}
                  value={formData.code}
                  onChange={(e) => handleChange('code', Number(e.target.value))}
                  className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                    errors.code ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                  }`}
                  placeholder="Ex: 100"
                />
                {errors.code && <p className="text-xs text-rose-400">{errors.code}</p>}
              </div>
            )}

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
                placeholder="Nome completo do representante"
              />
              {errors.name && <p className="text-xs text-rose-400">{errors.name}</p>}
            </div>

            {/* Email */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Email <span className="text-rose-400">*</span>
              </label>
              <input
                type="email"
                value={formData.email}
                onChange={(e) => handleChange('email', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.email ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="email@exemplo.com"
              />
              {errors.email && <p className="text-xs text-rose-400">{errors.email}</p>}
            </div>

            {/* Phone */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Telefone
              </label>
              <input
                type="text"
                value={formData.phone}
                onChange={(e) => handleChange('phone', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.phone ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="(11) 99999-9999"
              />
              {errors.phone && <p className="text-xs text-rose-400">{errors.phone}</p>}
            </div>

            {/* Company */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Empresa <span className="text-rose-400">*</span>
              </label>
              <input
                type="text"
                value={formData.company}
                onChange={(e) => handleChange('company', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.company ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="Nome da empresa"
              />
              {errors.company && <p className="text-xs text-rose-400">{errors.company}</p>}
            </div>

            {/* Region */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Região <span className="text-rose-400">*</span>
              </label>
              <input
                type="text"
                value={formData.region}
                onChange={(e) => handleChange('region', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.region ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="Ex: Sul, Sudeste"
              />
              {errors.region && <p className="text-xs text-rose-400">{errors.region}</p>}
            </div>

            {/* City */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Cidade <span className="text-rose-400">*</span>
              </label>
              <input
                type="text"
                value={formData.city}
                onChange={(e) => handleChange('city', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.city ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="Ex: São Paulo"
              />
              {errors.city && <p className="text-xs text-rose-400">{errors.city}</p>}
            </div>

            {/* Attendant */}
            <div className="space-y-2">
              <label className="block text-sm font-medium text-gray-300">
                Atendente
              </label>
              <input
                type="text"
                value={formData.attendant}
                onChange={(e) => handleChange('attendant', e.target.value)}
                className={`w-full px-4 py-2.5 bg-[#0f1115] border rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm transition-colors ${
                  errors.attendant ? 'border-rose-500 focus:ring-rose-500' : 'border-gray-700'
                }`}
                placeholder="Nome do atendente"
              />
              {errors.attendant && <p className="text-xs text-rose-400">{errors.attendant}</p>}
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
          </div>

          {/* Footer */}
          <div className="flex items-center justify-end gap-3 pt-6 border-t border-gray-800">
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

export default RepresentativeForm;
