import React, { useState } from 'react';
import { Copy, Check, ExternalLink, Key, Lock, Globe } from 'lucide-react';
import { useCredentials } from '../hooks/useSettings';
import type { ProgramCredential } from '@/shared/types';

const PasswordsView: React.FC = () => {
  const [copiedId, setCopiedId] = useState<string | null>(null);
  const { data: credentials = [], isLoading } = useCredentials();

  const copyToClipboard = (text: string, id: string) => {
    navigator.clipboard.writeText(text);
    setCopiedId(id);
    setTimeout(() => setCopiedId(null), 2000);
  };

  if (isLoading) {
    return (
      <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex items-center justify-center">
        <div className="text-gray-400">Carregando credenciais...</div>
      </div>
    );
  }

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      <div className="mb-6">
        <h1 className="text-2xl font-bold text-gray-100 flex items-center gap-3">
            <Lock className="text-emerald-500" />
            Acessos
        </h1>
        <p className="text-gray-500">Credenciais de acesso e ferramentas do Marketing</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-4 overflow-y-auto custom-scrollbar pb-4">
        {credentials.map((cred, idx) => (
            <div key={idx} className="bg-[#1a1d24] border border-gray-700/50 rounded-xl p-5 hover:border-emerald-500/30 transition-all group flex flex-col justify-between">
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
                    
                    <h3 className="font-bold text-lg text-white mb-3">{cred.name}</h3>
                    
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
                                        onClick={() => copyToClipboard(cred.user!, `user-${idx}`)}
                                        className="p-1.5 bg-gray-800 hover:bg-gray-700 text-gray-400 hover:text-white rounded border border-gray-700 transition-colors flex-shrink-0"
                                        title="Copiar Usuário"
                                    >
                                        {copiedId === `user-${idx}` ? <Check size={14} className="text-emerald-500" /> : <Copy size={14} />}
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
                                        onClick={() => copyToClipboard(cred.password!, `pass-${idx}`)}
                                        className="p-1.5 bg-gray-800 hover:bg-gray-700 text-gray-400 hover:text-white rounded border border-gray-700 transition-colors flex-shrink-0"
                                        title="Copiar Senha"
                                    >
                                        {copiedId === `pass-${idx}` ? <Check size={14} className="text-emerald-500" /> : <Copy size={14} />}
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
            </div>
        ))}
      </div>
    </div>
  );
};

export default PasswordsView;
