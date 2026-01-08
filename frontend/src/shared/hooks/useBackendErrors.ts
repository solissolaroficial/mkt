import { useToast } from './useToast';

interface BackendErrorResponse {
  message?: string;
  errors?: Record<string, string[]>;
  error?: string;
}

interface UseBackendErrorsOptions {
  setError?: (field: string, error: { type: string; message: string }) => void;
  onGlobalError?: (message: string) => void;
  fieldMapping?: Record<string, string>; // backend field → form field
}

export const useBackendErrors = (options: UseBackendErrorsOptions = {}) => {
  const { setError, onGlobalError, fieldMapping = {} } = options;
  const { error: showErrorToast } = useToast();

  const handleBackendErrors = (errorResponse: BackendErrorResponse) => {
    // Global error (sem field errors)
    if (!errorResponse.errors) {
      const message = errorResponse.message || errorResponse.error || 'Ocorreu um erro';
      if (onGlobalError) {
        onGlobalError(message);
      } else {
        showErrorToast(message);
      }
      return;
    }

    // Field errors
    if (setError) {
      Object.entries(errorResponse.errors).forEach(([field, messages]) => {
        const formField = fieldMapping[field] || field;
        setError(formField, {
          type: 'manual',
          message: messages[0], // Primeira mensagem de erro
        });
      });
    }
  };

  return { handleBackendErrors };
};
