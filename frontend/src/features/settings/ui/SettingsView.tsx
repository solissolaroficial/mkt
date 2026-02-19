import React, { useState, useEffect, useRef } from 'react';
import {
  User,
  Mail,
  Lock,
  Camera,
  Save,
  LogOut,
  Shield,
  BellRing,
  CheckCircle2,
  Archive,
  History,
  RefreshCcw,
  Clock,
  Target,
  AlertCircle,
  X,
  Upload,
  Trash2,
  Loader2,
  KeyRound
} from 'lucide-react';
import type { Notification, KpiCategory } from '@/shared/types';
import type { SettingsViewProps, SettingsSection, UserFormData } from '../types';
import { useChangePassword } from '../hooks/useChangePassword';
import { useUpdateProfile } from '../hooks/useUpdateProfile';
import { useGetProfile } from '../hooks/useGetProfile';
import { useUploadProfilePhoto } from '../hooks/useUploadProfilePhoto';
import { useRemoveProfilePhoto } from '../hooks/useRemoveProfilePhoto';
import { usePresignedURL } from '../hooks/usePresignedURL';
import { useAuth } from '@/features/auth/hooks/useAuth';
import { validatePassword, getPasswordStrengthColor, getPasswordStrengthLabel } from '../utils/passwordValidation';
import PasswordsView from './PasswordsView';

const SettingsView: React.FC<SettingsViewProps> = ({
    onLogout,
    notifications = [],
    onUpdateNotification,
    kpis,
    onUpdateKpiMeta
}) => {
  const [activeSection, setActiveSection] = useState<SettingsSection>('profile');
  const [showSuccess, setShowSuccess] = useState(false);
  const [passwordError, setPasswordError] = useState<string>('');
  const [passwordSuccess, setPasswordSuccess] = useState<string>('');
  const [showPasswordStrength, setShowPasswordStrength] = useState(false);
  const [profileError, setProfileError] = useState<string>('');
  const [photoError, setPhotoError] = useState<string>('');
  const [photoSuccess, setPhotoSuccess] = useState<string>('');
  const [selectedPhotoFile, setSelectedPhotoFile] = useState<File | null>(null);
  const [photoPreview, setPhotoPreview] = useState<string>('');
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Auth hooks
  const { user, setUser } = useAuth();

  // Password change hook
  const { changePassword, isLoading: isChangingPassword, error: changePasswordError, isSuccess: changePasswordSuccess, reset: resetPasswordMutation } = useChangePassword();

  // Update profile hook
  const { mutate: updateProfile, isPending: isUpdatingProfile, error: updateProfileError, isSuccess: updateProfileSuccess, reset: resetUpdateProfileMutation } = useUpdateProfile();

  // Upload profile photo hook
  const { mutate: uploadProfilePhoto, isPending: isUploadingPhoto, error: uploadPhotoError, isSuccess: uploadPhotoSuccess, reset: resetUploadPhotoMutation } = useUploadProfilePhoto();

  // Remove profile photo hook
  const { mutate: removeProfilePhoto, isPending: isRemovingPhoto, error: removePhotoError, isSuccess: removePhotoSuccess, reset: resetRemovePhotoMutation } = useRemoveProfilePhoto();

  // Get profile hook
  const { data: profileData, isLoading: isLoadingProfile } = useGetProfile();

  // Get presigned URL hook
  const { data: presignedURLData } = usePresignedURL();

  // Goal Editing State
  const [selectedGoalKpi, setSelectedGoalKpi] = useState<string>('');
  const [metaValues, setMetaValues] = useState<Record<string, number>>({});
  
  // Debounce timeout ref
  const debounceTimeoutRef = useRef<Record<string, NodeJS.Timeout>>({});

  // Atualizar selectedGoalKpi quando kpis mudar e estiver vazio
  useEffect(() => {
    if (kpis && kpis.length > 0 && !selectedGoalKpi) {
      const firstKpiId = kpis[0].id;
      setSelectedGoalKpi(firstKpiId);
      // Inicializar metaValues com os valores do primeiro KPI
      const initialValues: Record<string, number> = {};
      kpis[0].data.forEach(d => {
        if (d.meta !== null) {
          initialValues[d.month] = d.meta;
        }
      });
      setMetaValues(initialValues);
    }
  }, [kpis, selectedGoalKpi]);
  
  // Cleanup debounce timeouts on unmount
  useEffect(() => {
    return () => {
      Object.values(debounceTimeoutRef.current).forEach(timeout => clearTimeout(timeout));
    };
  }, []);

  // User Form State - initialized with real user data
  const [formData, setFormData] = useState<UserFormData>({
    name: user?.name || '',
    role: user?.role || '',
    email: user?.email || '',
    currentPassword: '',
    newPassword: '',
    confirmPassword: ''
  });

  // Update formData when user data changes
  useEffect(() => {
    if (user) {
      setFormData(prev => ({
        ...prev,
        name: user.name,
        role: user.role,
        email: user.email,
      }));
    }
  }, [user]);

  // Load user profile data if user is null (direct access to /settings)
  useEffect(() => {
    if (!user && !isLoadingProfile && profileData) {
      setUser(profileData);
    }
  }, [user, isLoadingProfile, profileData, setUser]);

  const handleSave = (e: React.FormEvent) => {
    e.preventDefault();
    setProfileError('');

    // Call API to update profile
    updateProfile({
      name: formData.name,
      email: formData.email,
      role: formData.role,
    });
  };

  // Handle profile update success/error
  useEffect(() => {
    if (updateProfileSuccess) {
      setShowSuccess(true);
      setTimeout(() => setShowSuccess(false), 3000);
      resetUpdateProfileMutation();
    }
  }, [updateProfileSuccess, resetUpdateProfileMutation]);

  useEffect(() => {
    if (updateProfileError) {
      // Extrair mensagem de erro específica do backend ou usar mensagem genérica
      const errorMessage = (updateProfileError as any).response?.data?.error ||
                         updateProfileError.message ||
                         'Erro ao atualizar perfil. Tente novamente.';
      setProfileError(errorMessage);
      resetUpdateProfileMutation();
    }
  }, [updateProfileError, resetUpdateProfileMutation]);

  const handlePasswordChange = (e: React.FormEvent) => {
    e.preventDefault();
    setPasswordError('');
    setPasswordSuccess('');
    
    // Validate passwords match
    if (formData.newPassword !== formData.confirmPassword) {
      setPasswordError('As senhas não coincidem');
      return;
    }
    
    // Validate password strength
    const validation = validatePassword(formData.newPassword);
    if (!validation.isValid) {
      setPasswordError(validation.errors[0]);
      return;
    }
    
    // Call API
    changePassword({
      currentPassword: formData.currentPassword,
      newPassword: formData.newPassword
    });
  };

  // Handle password change success/error
  useEffect(() => {
    if (changePasswordSuccess) {
      setPasswordSuccess('Senha alterada com sucesso!');
      setFormData(prev => ({
        ...prev,
        currentPassword: '',
        newPassword: '',
        confirmPassword: ''
      }));
      setTimeout(() => {
        setPasswordSuccess('');
        resetPasswordMutation();
      }, 3000);
    }
  }, [changePasswordSuccess, resetPasswordMutation]);

  useEffect(() => {
    if (changePasswordError) {
      // Extrair mensagem de erro específica do backend ou usar mensagem genérica
      const errorMessage = (changePasswordError as any).response?.data?.error ||
                         changePasswordError.message ||
                         'Erro ao alterar senha. Verifique sua senha atual e tente novamente.';
      setPasswordError(errorMessage);
      resetPasswordMutation();
    }
  }, [changePasswordError, resetPasswordMutation]);

  // Handle photo upload success/error
  useEffect(() => {
    if (uploadPhotoSuccess) {
      setPhotoSuccess('Foto de perfil atualizada com sucesso!');
      setTimeout(() => setPhotoSuccess(''), 3000);
      resetUploadPhotoMutation();
      // Nota: A invalidação de queries no hook useUploadProfilePhoto
      // irá recarregar os dados do usuário automaticamente
    }
  }, [uploadPhotoSuccess, resetUploadPhotoMutation]);

  useEffect(() => {
    if (uploadPhotoError) {
      const errorMessage = (uploadPhotoError as any).response?.data?.error ||
                         uploadPhotoError.message ||
                         'Erro ao fazer upload da foto. Tente novamente.';
      setPhotoError(errorMessage);
      resetUploadPhotoMutation();
    }
  }, [uploadPhotoError, resetUploadPhotoMutation]);

  // Handle photo remove success/error
  useEffect(() => {
    if (removePhotoSuccess) {
      setPhotoSuccess('Foto de perfil removida com sucesso!');
      setTimeout(() => setPhotoSuccess(''), 3000);
      resetRemovePhotoMutation();
    }
  }, [removePhotoSuccess, resetRemovePhotoMutation]);

  useEffect(() => {
    if (removePhotoError) {
      const errorMessage = (removePhotoError as any).response?.data?.error ||
                         removePhotoError.message ||
                         'Erro ao remover a foto. Tente novamente.';
      setPhotoError(errorMessage);
      resetRemovePhotoMutation();
    }
  }, [removePhotoError, resetRemovePhotoMutation]);

  const handlePhotoSelect = (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (!file) return;

    // Validate file type
    const allowedTypes = ['image/jpeg', 'image/jpg', 'image/png', 'image/gif'];
    if (!allowedTypes.includes(file.type)) {
      setPhotoError('Tipo de arquivo não suportado. Use JPG, PNG ou GIF.');
      return;
    }

    // Validate file size (5MB)
    const maxSize = 5 * 1024 * 1024;
    if (file.size > maxSize) {
      setPhotoError('Arquivo muito grande. Tamanho máximo: 5MB.');
      return;
    }

    setPhotoError('');
    setSelectedPhotoFile(file);
    
    // Create preview
    const reader = new FileReader();
    reader.onloadend = () => {
      setPhotoPreview(reader.result as string);
    };
    reader.readAsDataURL(file);

    // Upload photo automatically after selection
    uploadProfilePhoto(file);
  };

  const handlePhotoRemove = () => {
    removeProfilePhoto();
  };

  const handleUnarchive = (id: string) => {
    if (onUpdateNotification) {
        // Ao desarquivar, definimos read: false para que ela volte ao topo da lista e notifique o usuário
        const updated = notifications.map(n => n.id === id ? { ...n, archived: false, read: false } : n);
        onUpdateNotification(updated);
    }
  };

  const archivedNotifications = notifications.filter(n => n.archived);

  // Filter KPI for goals
  const currentKpiForGoals = kpis?.find(k => k.id === selectedGoalKpi);

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 max-w-5xl mx-auto pb-10">
      
      {/* Header */}
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-100 mb-2">Configurações da Conta</h1>
        <p className="text-gray-500">Gerencie suas informações pessoais e segurança.</p>
      </div>

      <div className="flex flex-col lg:flex-row gap-8">
        
        {/* Sidebar Navigation (Local) */}
        <div className="w-full lg:w-64 flex-shrink-0 space-y-2">
            <button 
                onClick={() => setActiveSection('profile')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'profile' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <User size={18} /> Perfil Público
            </button>
            <button 
                onClick={() => setActiveSection('security')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'security' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <Shield size={18} /> Segurança
            </button>
            <button 
                onClick={() => setActiveSection('history')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'history' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <History size={18} /> Histórico de Notificações
            </button>
            {kpis && onUpdateKpiMeta && (
                <button 
                    onClick={() => setActiveSection('goals')}
                    className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'goals' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
                >
                    <Target size={18} /> Metas do Sistema
                </button>
            )}
            <button 
                onClick={() => setActiveSection('passwords')}
                className={`w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 transition-colors ${activeSection === 'passwords' ? 'bg-[#1e5144]/10 text-emerald-400 border border-[#1e5144]/20' : 'text-gray-400 hover:bg-[#1a1d24] hover:text-gray-200'}`}
            >
                <KeyRound size={18} /> Acessos
            </button>
            
            <div className="pt-8 mt-8 border-t border-gray-800">
                <button 
                    onClick={onLogout}
                    className="w-full text-left px-4 py-3 rounded-xl font-medium flex items-center gap-3 text-rose-400 hover:bg-rose-500/10 hover:text-rose-300 transition-colors border border-transparent hover:border-rose-500/20"
                >
                    <LogOut size={18} /> Sair do Sistema
                </button>
            </div>
        </div>

        {/* Main Content Form */}
        <div className="flex-grow">
            {activeSection === 'profile' && (
                <form onSubmit={handleSave} className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden">
                        <div className="p-8 space-y-8">
                            {/* Error Message */}
                            {profileError && (
                                <div className="flex items-start gap-2 p-3 bg-rose-500/10 border border-rose-500/20 rounded-lg animate-in fade-in slide-in-from-top-2">
                                    <AlertCircle size={18} className="text-rose-400 flex-shrink-0 mt-0.5" />
                                    <div className="flex-grow">
                                        <p className="text-sm text-rose-300">{profileError}</p>
                                    </div>
                                    <button
                                        onClick={() => setProfileError('')}
                                        className="text-rose-400 hover:text-rose-300"
                                    >
                                        <X size={16} />
                                    </button>
                                </div>
                            )}

                            {/* Avatar Section */}
                            <div className="flex items-center gap-6 pb-8 border-b border-gray-800">
                                <div className="relative group">
                                    <div className="w-24 h-24 rounded-full bg-gradient-to-tr from-gray-700 to-gray-600 border-4 border-[#1a1d24] flex items-center justify-center text-3xl font-bold text-white shadow-lg overflow-hidden">
                                        {photoPreview ? (
                                            <img
                                                src={photoPreview}
                                                alt="Profile"
                                                className="w-full h-full object-cover"
                                            />
                                        ) : presignedURLData?.url ? (
                                            <img
                                                src={presignedURLData.url}
                                                alt="Profile"
                                                className="w-full h-full object-cover"
                                            />
                                        ) : (
                                            formData.name.substring(0, 2).toUpperCase()
                                        )}
                                    </div>
                                    <div className="absolute inset-0 bg-black/50 rounded-full flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity cursor-pointer">
                                        <Camera size={24} className="text-white" />
                                    </div>
                                </div>
                                <div>
                                    <h3 className="text-lg font-bold text-gray-100">Sua Foto</h3>
                                    <p className="text-sm text-gray-500 mb-3">Isso será exibido no seu perfil.</p>
                                    <div className="flex gap-3">
                                        <input
                                            ref={fileInputRef}
                                            type="file"
                                            accept="image/jpeg,image/jpg,image/png,image/gif"
                                            onChange={handlePhotoSelect}
                                            className="hidden"
                                        />
                                        <button
                                            type="button"
                                            onClick={() => fileInputRef.current?.click()}
                                            className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white text-xs font-semibold rounded-lg transition-colors flex items-center gap-2"
                                        >
                                            {isUploadingPhoto ? (
                                                <>
                                                    <Loader2 size={14} className="animate-spin" />
                                                    Enviando...
                                                </>
                                            ) : (
                                                <>
                                                    <Upload size={14} />
                                                    Alterar
                                                </>
                                            )}
                                        </button>
                                        {(photoPreview || presignedURLData?.url) && (
                                            <button
                                                type="button"
                                                onClick={handlePhotoRemove}
                                                disabled={isRemovingPhoto}
                                                className="px-4 py-2 bg-gray-800 hover:bg-gray-700 text-gray-300 text-xs font-semibold rounded-lg transition-colors flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
                                            >
                                                {isRemovingPhoto ? (
                                                    <Loader2 size={14} className="animate-spin" />
                                                ) : (
                                                    <Trash2 size={14} />
                                                )}
                                                {isRemovingPhoto ? 'Removendo...' : 'Remover'}
                                            </button>
                                        )}
                                    </div>
                                    {/* Photo Error/Success Messages */}
                                    {photoError && (
                                        <div className="mt-3 text-sm text-rose-400 flex items-center gap-1">
                                            <AlertCircle size={14} />
                                            {photoError}
                                        </div>
                                    )}
                                    {photoSuccess && (
                                        <div className="mt-3 text-sm text-emerald-400 flex items-center gap-1">
                                            <CheckCircle2 size={14} />
                                            {photoSuccess}
                                        </div>
                                    )}
                                </div>
                            </div>

                            {/* Form Fields */}
                            <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Nome Completo</label>
                                    <div className="relative">
                                        <User size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input
                                            type="text"
                                            value={formData.name}
                                            onChange={(e) => setFormData({...formData, name: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Cargo</label>
                                    <input
                                        type="text"
                                        value={formData.role}
                                        disabled
                                        className="w-full px-4 py-2.5 bg-[#0f1115]/50 border border-gray-700/50 rounded-xl text-gray-500 cursor-not-allowed text-sm"
                                    />
                                </div>
                                <div className="md:col-span-2">
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">E-mail</label>
                                    <div className="relative">
                                        <Mail size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input
                                            type="email"
                                            value={formData.email}
                                            onChange={(e) => setFormData({...formData, email: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                        />
                                    </div>
                                </div>
                            </div>
                        </div>
                        {/* Footer Actions */}
                        <div className="p-6 bg-[#20232b] border-t border-gray-800 flex justify-between items-center">
                            <div>
                                {showSuccess && (
                                    <span className="flex items-center gap-2 text-emerald-400 text-sm font-medium animate-in fade-in slide-in-from-left-2">
                                        <CheckCircle2 size={16} /> Alterações salvas com sucesso!
                                    </span>
                                )}
                            </div>
                            <button
                                type="submit"
                                disabled={isUpdatingProfile}
                                className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white font-semibold rounded-xl flex items-center gap-2 shadow-lg shadow-[#1e5144]/20 transition-all disabled:opacity-70 disabled:cursor-not-allowed"
                            >
                                {isUpdatingProfile ? (
                                    <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                                ) : (
                                    <Save size={18} />
                                )}
                                Salvar Alterações
                            </button>
                        </div>
                </form>
            )}

            {activeSection === 'security' && (
                <form onSubmit={handlePasswordChange} className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden">
                        <div className="p-8 space-y-8">
                            <div className="pb-6 border-b border-gray-800">
                                <h3 className="text-lg font-bold text-gray-100">Alterar Senha</h3>
                                <p className="text-sm text-gray-500">Mantenha sua conta segura alterando sua senha regularmente.</p>
                            </div>

                            <div className="space-y-4 max-w-md">
                                {/* Error Message */}
                                {passwordError && (
                                    <div className="flex items-start gap-2 p-3 bg-rose-500/10 border border-rose-500/20 rounded-lg animate-in fade-in slide-in-from-top-2">
                                        <AlertCircle size={18} className="text-rose-400 flex-shrink-0 mt-0.5" />
                                        <div className="flex-grow">
                                            <p className="text-sm text-rose-300">{passwordError}</p>
                                        </div>
                                        <button
                                            onClick={() => setPasswordError('')}
                                            className="text-rose-400 hover:text-rose-300"
                                        >
                                            <X size={16} />
                                        </button>
                                    </div>
                                )}
                                
                                {/* Success Message */}
                                {passwordSuccess && (
                                    <div className="flex items-start gap-2 p-3 bg-emerald-500/10 border border-emerald-500/20 rounded-lg animate-in fade-in slide-in-from-top-2">
                                        <CheckCircle2 size={18} className="text-emerald-400 flex-shrink-0 mt-0.5" />
                                        <div className="flex-grow">
                                            <p className="text-sm text-emerald-300">{passwordSuccess}</p>
                                        </div>
                                        <button
                                            onClick={() => setPasswordSuccess('')}
                                            className="text-emerald-400 hover:text-emerald-300"
                                        >
                                            <X size={16} />
                                        </button>
                                    </div>
                                )}
                                
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Senha Atual</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input
                                            type="password"
                                            value={formData.currentPassword}
                                            onChange={(e) => setFormData({...formData, currentPassword: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Nova Senha</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input
                                            type="password"
                                            value={formData.newPassword}
                                            onChange={(e) => {
                                                setFormData({...formData, newPassword: e.target.value});
                                                setShowPasswordStrength(e.target.value.length > 0);
                                            }}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                    {/* Password Strength Indicator */}
                                    {showPasswordStrength && formData.newPassword.length > 0 && (
                                        <div className="mt-2">
                                            <div className="flex items-center gap-2 mb-1">
                                                <div className="flex-grow h-1.5 bg-gray-700 rounded-full overflow-hidden">
                                                    <div
                                                        className="h-full transition-all duration-300"
                                                        style={{
                                                            width: formData.newPassword.length >= 8 ? '100%' : '50%',
                                                            backgroundColor: getPasswordStrengthColor(validatePassword(formData.newPassword).strength)
                                                        }}
                                                    />
                                                </div>
                                                <span
                                                    className="text-xs font-semibold"
                                                    style={{ color: getPasswordStrengthColor(validatePassword(formData.newPassword).strength) }}
                                                >
                                                    {getPasswordStrengthLabel(validatePassword(formData.newPassword).strength)}
                                                </span>
                                            </div>
                                            <p className="text-[10px] text-gray-500">
                                                Mínimo 8 caracteres, incluindo 3 dos 4 tipos: maiúsculas, minúsculas, números ou caracteres especiais
                                            </p>
                                        </div>
                                    )}
                                </div>
                                <div>
                                    <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Confirmar Nova Senha</label>
                                    <div className="relative">
                                        <Lock size={16} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" />
                                        <input
                                            type="password"
                                            value={formData.confirmPassword}
                                            onChange={(e) => setFormData({...formData, confirmPassword: e.target.value})}
                                            className="w-full pl-10 pr-4 py-2.5 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                                            placeholder="••••••••"
                                        />
                                    </div>
                                    {/* Password Match Indicator */}
                                    {formData.confirmPassword.length > 0 && (
                                        <div className="mt-1 flex items-center gap-1.5">
                                            {formData.newPassword === formData.confirmPassword ? (
                                                <>
                                                    <CheckCircle2 size={14} className="text-emerald-400" />
                                                    <span className="text-[10px] text-emerald-400">As senhas coincidem</span>
                                                </>
                                            ) : (
                                                <>
                                                    <AlertCircle size={14} className="text-rose-400" />
                                                    <span className="text-[10px] text-rose-400">As senhas não coincidem</span>
                                                </>
                                            )}
                                        </div>
                                    )}
                                </div>
                            </div>
                        </div>
                        <div className="p-6 bg-[#20232b] border-t border-gray-800 flex justify-between items-center">
                            <div>
                                {passwordSuccess && (
                                    <span className="flex items-center gap-2 text-emerald-400 text-sm font-medium animate-in fade-in slide-in-from-left-2">
                                        <CheckCircle2 size={16} /> {passwordSuccess}
                                    </span>
                                )}
                            </div>
                            <button
                                type="submit"
                                disabled={isChangingPassword || !formData.currentPassword || !formData.newPassword || !formData.confirmPassword}
                                className="px-6 py-2.5 bg-[#1e5144] hover:bg-[#163c32] text-white font-semibold rounded-xl flex items-center gap-2 shadow-lg shadow-[#1e5144]/20 transition-all disabled:opacity-50 disabled:cursor-not-allowed"
                            >
                                {isChangingPassword ? (
                                    <div className="w-4 h-4 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
                                ) : (
                                    <Save size={18} />
                                )}
                                Salvar Senha
                            </button>
                        </div>
                </form>
            )}

            {activeSection === 'history' && (
                <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden h-full flex flex-col">
                    <div className="p-6 border-b border-gray-800 bg-[#20232b] flex justify-between items-center">
                        <div>
                            <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2">
                                <Archive size={20} className="text-gray-400" />
                                Notificações Arquivadas
                            </h3>
                            <p className="text-sm text-gray-500">Histórico de alertas que você arquivou.</p>
                        </div>
                    </div>
                    
                    <div className="flex-grow overflow-y-auto custom-scrollbar">
                        {archivedNotifications.length === 0 ? (
                            <div className="p-12 text-center text-gray-500">
                                <Archive size={48} className="mx-auto mb-4 opacity-20" />
                                <p className="text-lg font-medium">Histórico vazio</p>
                                <p className="text-sm">Você não tem notificações arquivadas.</p>
                            </div>
                        ) : (
                            <div className="divide-y divide-gray-800">
                                {archivedNotifications.map(notification => (
                                    <div key={notification.id} className="p-4 hover:bg-[#20232b] transition-colors flex gap-4">
                                        <div className="w-10 h-10 rounded-full bg-gray-800 border border-gray-700 flex items-center justify-center flex-shrink-0 text-gray-400">
                                            {notification.notification_type === 'mention' && <User size={18} />}
                                            {notification.notification_type === 'deadline' && <Clock size={18} />}
                                            {notification.notification_type === 'system' && <BellRing size={18} />}
                                        </div>
                                        <div className="flex-grow">
                                            <div className="flex justify-between items-start">
                                                <h4 className="text-sm font-semibold text-gray-200">{notification.title}</h4>
                                                <span className="text-[10px] text-gray-500">{notification.timestamp}</span>
                                            </div>
                                            <p className="text-sm text-gray-400 mt-1">{notification.message}</p>
                                        </div>
                                        <button 
                                            onClick={() => handleUnarchive(notification.id)}
                                            className="p-2 text-gray-500 hover:text-emerald-400 hover:bg-emerald-500/10 rounded-lg transition-colors self-center"
                                            title="Desarquivar"
                                        >
                                            <RefreshCcw size={18} />
                                        </button>
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>
                </div>
            )}

            {activeSection === 'goals' && currentKpiForGoals && onUpdateKpiMeta && (
                <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden h-full flex flex-col">
                    <div className="p-6 border-b border-gray-800 bg-[#20232b]">
                        <h3 className="text-lg font-bold text-gray-100 flex items-center gap-2 mb-4">
                            <Target size={20} className="text-emerald-400" />
                            Ajuste de Metas
                        </h3>
                        
                        <div>
                            <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">Selecione o Indicador (KPI)</label>
                            <select
                                value={selectedGoalKpi}
                                onChange={(e) => {
                                    const newKpiId = e.target.value;
                                    setSelectedGoalKpi(newKpiId);
                                    // Inicializar metaValues com os valores do novo KPI
                                    const newKpi = kpis?.find(k => k.id === newKpiId);
                                    if (newKpi) {
                                        const initialValues: Record<string, number> = {};
                                        newKpi.data.forEach(d => {
                                            if (d.meta !== null) {
                                                initialValues[d.month] = d.meta;
                                            }
                                        });
                                        setMetaValues(initialValues);
                                    }
                                }}
                                className="w-full md:w-1/2 bg-[#0f1115] border border-gray-700 rounded-lg px-3 py-2 text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm"
                            >
                                {kpis?.map(k => (
                                    <option key={k.id} value={k.id}>{k.title}</option>
                                ))}
                            </select>
                        </div>
                    </div>

                    <div className="p-6 overflow-y-auto custom-scrollbar flex-grow">
                        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                            {currentKpiForGoals.data.map((monthlyData, index) => (
                                <div key={monthlyData.month} className="bg-[#0f1115] p-4 rounded-xl border border-gray-800 hover:border-gray-700 transition-colors">
                                    <div className="flex justify-between items-center mb-2">
                                        <span className="font-bold text-gray-300">{monthlyData.month}</span>
                                        <span className="text-xs text-gray-500">Mês {index + 1}</span>
                                    </div>
                                    <div className="relative">
                                        <label className="block text-[10px] text-gray-500 uppercase font-bold mb-1">Meta Definida</label>
                                        <input
                                            type="number"
                                            value={metaValues[monthlyData.month] !== undefined ? metaValues[monthlyData.month] : (monthlyData.meta !== null ? monthlyData.meta : '')}
                                            onChange={(e) => {
                                                const val = parseFloat(e.target.value);
                                                const currentYear = new Date().getFullYear();
                                                const monthKey = monthlyData.month;
                                                
                                                // Atualizar estado local imediatamente para o input funcionar
                                                setMetaValues(prev => ({ ...prev, [monthKey]: isNaN(val) ? 0 : val }));
                                                
                                                // Limpar timeout anterior se existir
                                                if (debounceTimeoutRef.current[monthKey]) {
                                                    clearTimeout(debounceTimeoutRef.current[monthKey]);
                                                }
                                                
                                                // Debounce: aguardar 800ms antes de enviar ao backend
                                                debounceTimeoutRef.current[monthKey] = setTimeout(() => {
                                                    onUpdateKpiMeta(currentKpiForGoals.id, monthKey, currentYear, isNaN(val) ? 0 : val);
                                                }, 800);
                                            }}
                                            className="w-full bg-[#1a1d24] border border-gray-700 rounded-lg px-3 py-2 text-white focus:outline-none focus:ring-1 focus:ring-[#1e5144] font-mono text-sm"
                                            placeholder="Definir meta"
                                        />
                                        <span className="absolute right-3 top-[26px] text-gray-500 text-xs pointer-events-none">
                                            {currentKpiForGoals.unit === 'currency' ? 'R$' : currentKpiForGoals.unit === 'percent' ? '%' : 'un'}
                                        </span>
                                    </div>
                                </div>
                            ))}
                        </div>
                    </div>
                    
                    <div className="p-4 border-t border-gray-800 bg-[#20232b] text-center">
                        <p className="text-xs text-gray-500">
                            As alterações são salvas automaticamente no sistema.
                        </p>
                    </div>
                </div>
            )}

            {activeSection === 'passwords' && <PasswordsView />}
        </div>
      </div>
    </div>
  );
};

export default SettingsView;
