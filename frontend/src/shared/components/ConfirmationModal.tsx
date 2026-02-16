import React from 'react';
import { X } from 'lucide-react';

interface ConfirmationModalProps {
  isOpen: boolean;
  onClose: () => void;
  onConfirm: () => void;
  title: string;
  message: string;
  isPending?: boolean;
  type?: 'danger' | 'warning' | 'info';
}

const ConfirmationModal: React.FC<ConfirmationModalProps> = ({
  isOpen,
  onClose,
  onConfirm,
  title,
  message,
  isPending = false,
  type = 'danger'
}) => {
  if (!isOpen) return null;

  const typeConfig = {
    danger: {
      buttonColor: 'bg-rose-600 hover:bg-rose-700',
      iconColor: 'text-rose-400'
    },
    warning: {
      buttonColor: 'bg-amber-600 hover:bg-amber-700',
      iconColor: 'text-amber-400'
    },
    info: {
      buttonColor: 'bg-blue-600 hover:bg-blue-700',
      iconColor: 'text-blue-400'
    }
  };

  const config = typeConfig[type];

  return (
    <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm" onClick={onClose}>
      <div
        className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden animate-in zoom-in-95 duration-200"
        onClick={e => e.stopPropagation()}
      >
        <div className="p-6 border-b border-gray-800 flex justify-between items-center bg-[#20232b]">
          <h3 className="text-lg font-bold text-gray-100">{title}</h3>
          <button onClick={onClose} className="text-gray-500 hover:text-white p-1 hover:bg-gray-800 rounded-full transition-colors">
            <X size={20} />
          </button>
        </div>
        <div className="p-6 bg-[#0f1115]">
          <p className="text-sm text-gray-300">{message}</p>
        </div>
        <div className="p-6 flex justify-end gap-3 bg-[#20232b]">
          <button
            onClick={onClose}
            className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
          >
            Cancelar
          </button>
          <button
            onClick={onConfirm}
            disabled={isPending}
            className={`px-4 py-2 text-white rounded-lg text-sm font-medium transition-colors ${config.buttonColor} disabled:opacity-50`}
          >
            {isPending ? 'Processando...' : 'Confirmar'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default ConfirmationModal;
