import React, { Component, ReactNode } from 'react';

interface ErrorBoundaryProps {
  children: ReactNode;
  fallback?: ReactNode;
}

interface ErrorBoundaryState {
  hasError: boolean;
  error: Error | null;
}

/**
 * ErrorBoundary component to catch JavaScript errors in component trees
 * Provides a fallback UI when an error occurs
 */
export class ErrorBoundary extends Component<ErrorBoundaryProps, ErrorBoundaryState> {
  constructor(props: ErrorBoundaryProps) {
    super(props);
    this.state = { hasError: false, error: null };
  }

  static getDerivedStateFromError(error: Error): ErrorBoundaryState {
    return { hasError: true, error };
  }

  componentDidCatch(error: Error, errorInfo: React.ErrorInfo) {
    console.error('ErrorBoundary capturou um erro:', error, errorInfo);
    console.error('Error Info:', {
      componentStack: errorInfo.componentStack,
    });
    
    // TODO: Enviar para serviço de logging (Sentry, LogRocket, etc.)
    // logErrorToService(error, errorInfo);
  }

  handleReset = () => {
    this.setState({ hasError: false, error: null });
  };

  render() {
    if (this.state.hasError) {
      return this.props.fallback || (
        <div className="min-h-screen bg-[#0f1115] flex items-center justify-center p-4">
          <div className="max-w-md w-full bg-[#1a1d24] rounded-2xl p-8 border border-red-500/30 shadow-2xl">
            <div className="text-center">
              <div className="w-16 h-16 bg-red-500/10 rounded-full flex items-center justify-center mx-auto mb-4">
                <svg className="w-8 h-8 text-red-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 2.502-3.215V8.108c0-1.548-1.962-3.215-2.502-3.215H4.938c-1.54 0-2.502 1.667-2.502 3.215v10.677c0 1.548 1.962 3.215 2.502 3.215h13.856c1.54 0 2.502-1.667 2.502-3.215V8.108c0-1.548-1.962-3.215-2.502-3.215z" />
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15h.01" />
                </svg>
              </div>
              <h2 className="text-xl font-bold text-white mb-2">Ops! Algo deu errado</h2>
              <p className="text-gray-400 mb-6">
                Ocorreu um erro inesperado. Tente recarregar a página ou voltar para a página anterior.
              </p>
              
              {this.state.error && (
                <details className="mb-4 text-left">
                  <summary className="cursor-pointer text-sm text-gray-500 hover:text-gray-400 mb-2">
                    Ver detalhes do erro
                  </summary>
                  <div className="mt-2 p-4 bg-[#0f1115] rounded-lg text-xs text-red-400 font-mono overflow-auto max-h-40">
                    <p className="font-bold mb-1">{this.state.error.name}</p>
                    <p className="break-words">{this.state.error.message}</p>
                    {this.state.error.stack && (
                      <pre className="mt-2 text-gray-500 whitespace-pre-wrap">
                        {this.state.error.stack}
                      </pre>
                    )}
                  </div>
                </details>
              )}
              
              <div className="flex flex-col gap-3">
                <button
                  onClick={this.handleReset}
                  className="w-full px-4 py-3 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg font-medium transition-colors"
                >
                  Tentar Novamente
                </button>
                <button
                  onClick={() => window.location.href = '/kpis'}
                  className="w-full px-4 py-3 bg-[#1a1d24] border border-gray-700 hover:border-gray-600 text-gray-300 rounded-lg font-medium transition-colors"
                >
                  Voltar para KPIs
                </button>
              </div>
            </div>
          </div>
        </div>
      );
    }

    return this.props.children;
  }
}

export default ErrorBoundary;
