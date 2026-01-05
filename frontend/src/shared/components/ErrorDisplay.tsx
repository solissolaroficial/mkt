import React from 'react';
import { AlertCircle, RefreshCw } from 'lucide-react';

interface ErrorDisplayProps {
  error: unknown;
  onRetry?: () => void;
}

export const ErrorDisplay: React.FC<ErrorDisplayProps> = ({ error, onRetry }) => {
  const getErrorMessage = (err: unknown): string => {
    if (err instanceof Error) {
      // Erros de rede
      if (err.message.includes('Network Error') || err.message.includes('ERR_NETWORK')) {
        return 'Não foi possível conectar ao servidor. Verifique sua conexão.';
      }
      // Erros de autenticação
      if (err.message.includes('401') || err.message.includes('Unauthorized')) {
        return 'Sessão expirada. Por favor, faça login novamente.';
      }
      // Erros de servidor
      if (err.message.includes('500') || err.message.includes('Internal Server Error')) {
        return 'Erro interno do servidor. Tente novamente em alguns minutos.';
      }
      // Erro genérico
      return err.message;
    }
    return 'Ocorreu um erro inesperado.';
  };

  return (
    <div className="flex flex-col items-center justify-center h-full p-6">
      <div className="bg-red-900/20 border border-red-800 rounded-lg p-6 max-w-md text-center">
        <AlertCircle className="w-12 h-12 text-red-400 mx-auto mb-4" />
        <h3 className="text-lg font-semibold text-red-100 mb-2">
          Erro ao carregar dados
        </h3>
        <p className="text-red-300 mb-4">
          {getErrorMessage(error)}
        </p>
        {onRetry && (
          <button
            onClick={onRetry}
            className="inline-flex items-center gap-2 bg-red-700 text-white px-4 py-2 rounded-lg hover:bg-red-600 transition-colors"
          >
            <RefreshCw size={16} />
            Tentar novamente
          </button>
        )}
      </div>
    </div>
  );
};
