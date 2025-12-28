import React, { useState } from 'react';
import { Eye, EyeOff, Lock, Mail, ArrowRight } from 'lucide-react';
import { useLogin } from '../hooks/useLogin';

const LoginView: React.FC = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPassword, setShowPassword] = useState(false);

  const { mutate: login, isPending, error } = useLogin();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    login({ email, password });
  };

  return (
    <div className="min-h-screen bg-[#0f1115] flex items-center justify-center p-4">
      <div className="w-full max-w-md">
        {/* Logo / Brand */}
        <div className="flex justify-center mb-8">
          <div className="flex items-center gap-3">
            <div className="w-12 h-12 bg-gradient-to-br from-[#1e5144] to-emerald-600 rounded-xl flex items-center justify-center shadow-lg shadow-emerald-900/50">
              <span className="font-bold text-white text-2xl">S</span>
            </div>
            <div>
              <h1 className="font-bold text-3xl tracking-tight text-white">
                Solis <span className="text-emerald-500">Hub</span>
              </h1>
            </div>
          </div>
        </div>

        <div className="bg-[#1a1d24] border border-gray-800 rounded-2xl shadow-2xl p-8 animate-in fade-in slide-in-from-bottom-4 duration-500">
          <h2 className="text-2xl font-bold text-gray-100 mb-2 text-center">
            Bem-vindo de volta
          </h2>
          <p className="text-gray-500 text-center mb-8 text-sm">
            Insira suas credenciais para acessar o painel.
          </p>

          {/* Mensagem de erro */}
          {error && (
            <div className="mb-4 p-3 bg-red-500/10 border border-red-500/50 rounded-lg text-red-400 text-sm">
              {error.message || 'Erro ao fazer login. Verifique suas credenciais.'}
            </div>
          )}

          <form onSubmit={handleSubmit} className="space-y-5">
            <div>
              <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">
                E-mail Corporativo
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                  <Mail size={18} />
                </div>
                <input
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  className="w-full pl-10 pr-4 py-3 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all placeholder-gray-600"
                  placeholder="seu.nome@solis.com.br"
                  required
                  disabled={isPending}
                />
              </div>
            </div>

            <div>
              <label className="block text-xs font-semibold text-gray-400 uppercase tracking-wider mb-2">
                Senha
              </label>
              <div className="relative">
                <div className="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none text-gray-500">
                  <Lock size={18} />
                </div>
                <input
                  type={showPassword ? 'text' : 'password'}
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  className="w-full pl-10 pr-10 py-3 bg-[#0f1115] border border-gray-700 rounded-xl text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] focus:border-transparent transition-all placeholder-gray-600"
                  placeholder="••••••••"
                  required
                  disabled={isPending}
                />
                <button
                  type="button"
                  onClick={() => setShowPassword(!showPassword)}
                  className="absolute inset-y-0 right-0 pr-3 flex items-center text-gray-500 hover:text-gray-300 transition-colors"
                  disabled={isPending}
                >
                  {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
                </button>
              </div>
            </div>

            <div className="flex items-center justify-between text-sm">
              <label className="flex items-center gap-2 cursor-pointer group">
                <input
                  type="checkbox"
                  className="w-4 h-4 rounded bg-[#0f1115] border-gray-700 text-[#1e5144] focus:ring-offset-0 focus:ring-[#1e5144]"
                  disabled={isPending}
                />
                <span className="text-gray-500 group-hover:text-gray-400 transition-colors">
                  Lembrar-me
                </span>
              </label>
              <a
                href="#"
                className="text-[#1e5144] hover:text-emerald-400 font-medium transition-colors"
              >
                Esqueceu a senha?
              </a>
            </div>

            <button
              type="submit"
              disabled={isPending}
              className="w-full bg-[#1e5144] hover:bg-[#163c32] text-white font-semibold py-3 rounded-xl transition-all duration-300 transform hover:scale-[1.02] flex items-center justify-center gap-2 shadow-lg shadow-[#1e5144]/20 disabled:opacity-70 disabled:cursor-not-allowed disabled:hover:scale-100"
            >
              {isPending ? (
                <div className="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"></div>
              ) : (
                <>
                  Entrar no Sistema <ArrowRight size={18} />
                </>
              )}
            </button>
          </form>
        </div>

        <p className="text-center text-gray-600 text-xs mt-8">
          &copy; 2025 Solis Solar. Todos os direitos reservados.
        </p>
      </div>
    </div>
  );
};

export default LoginView;