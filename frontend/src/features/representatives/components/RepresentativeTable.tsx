import React from 'react';
import { Search, Edit2, Trash2, Eye, Plus } from 'lucide-react';
import type { Representative, RepresentativeTableData } from '../types';

interface RepresentativeTableProps {
  data: RepresentativeTableData;
  isLoading: boolean;
  error: unknown;
  onRefetch: () => void;
  onCreate: () => void;
  onEdit: (representative: Representative) => void;
  onView: (representative: Representative) => void;
  onDelete: (uuid: string) => void;
  onPageChange: (page: number) => void;
  onLimitChange: (limit: number) => void;
  onSearch: (term: string) => void;
  onSortChange: (sortBy: string, sortOrder: 'asc' | 'desc') => void;
  currentFilters?: {
    page?: number;
    limit?: number;
    sortBy?: string;
    sortOrder?: 'asc' | 'desc';
    name?: string;
    company?: string;
    region?: string;
    city?: string;
    active?: boolean;
  };
}

const RepresentativeTable: React.FC<RepresentativeTableProps> = ({
  data,
  isLoading,
  error,
  onRefetch,
  onCreate,
  onEdit,
  onView,
  onDelete,
  onPageChange,
  onLimitChange,
  onSearch,
  onSortChange,
  currentFilters,
}) => {
  const [searchTerm, setSearchTerm] = React.useState('');
  const [localPage, setLocalPage] = React.useState(currentFilters?.page || 1);
  const [localLimit, setLocalLimit] = React.useState(currentFilters?.limit || 10);

  React.useEffect(() => {
    setLocalPage(currentFilters?.page || 1);
    setLocalLimit(currentFilters?.limit || 10);
  }, [currentFilters]);

  const handleSearch = (value: string) => {
    setSearchTerm(value);
    onSearch(value);
    setLocalPage(1);
  };

  const handleSort = (field: string) => {
    const currentSort = currentFilters?.sortBy;
    const currentOrder = currentFilters?.sortOrder || 'asc';
    
    if (currentSort === field) {
      onSortChange(field, currentOrder === 'asc' ? 'desc' : 'asc');
    } else {
      onSortChange(field, 'asc');
    }
  };

  const handlePageChange = (newPage: number) => {
    setLocalPage(newPage);
    onPageChange(newPage);
  };

  const handleLimitChange = (newLimit: number) => {
    setLocalLimit(newLimit);
    onLimitChange(newLimit);
    setLocalPage(1);
  };

  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-gray-400">Carregando representantes...</div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="flex flex-col items-center justify-center h-full gap-4">
        <div className="text-rose-400">Erro ao carregar representantes</div>
        <button
          onClick={onRefetch}
          className="px-4 py-2 bg-[#1e5144] hover:bg-[#163c32] text-white rounded-lg"
        >
          Tentar novamente
        </button>
      </div>
    );
  }

  const getSortIcon = (field: string) => {
    if (currentFilters?.sortBy !== field) return null;
    return currentFilters?.sortOrder === 'asc' ? '↑' : '↓';
  };

  return (
    <div className="animate-in fade-in slide-in-from-bottom-4 duration-500 h-full flex flex-col">
      {/* Header */}
      <div className="flex flex-col md:flex-row justify-between items-start md:items-center gap-4 mb-6">
        <div className="flex items-center gap-4">
          <h1 className="text-2xl font-bold text-gray-100">Representantes</h1>
          <p className="text-gray-500 text-sm">Gerencie seus representantes comerciais</p>
        </div>
        <div className="flex gap-3 w-full md:w-auto">
          <div className="relative flex-grow md:flex-grow-0">
            <Search className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-500" size={16} />
            <input
              type="text"
              placeholder="Buscar representante..."
              value={searchTerm}
              onChange={(e) => handleSearch(e.target.value)}
              className="w-full md:w-64 pl-10 pr-4 py-2 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 focus:outline-none focus:ring-2 focus:ring-[#1e5144] text-sm placeholder-gray-600"
            />
          </div>
          <button
            onClick={onCreate}
            className="bg-[#1e5144] hover:bg-[#163c32] text-white px-4 py-2 rounded-lg text-sm flex items-center gap-2 transition-colors shadow-lg"
          >
            <Plus size={16} /> Novo Representante
          </button>
        </div>
      </div>

      {/* Table */}
      <div className="bg-[#1a1d24] rounded-2xl border border-white/5 shadow-xl overflow-hidden flex-grow flex flex-col">
        <div className="overflow-auto custom-scrollbar flex-grow">
          <table className="w-full text-sm text-left whitespace-nowrap border-collapse">
            <thead className="bg-[#15171c] text-gray-400 uppercase font-semibold sticky top-0 z-20 shadow-md">
              <tr>
                <th className="px-4 py-3 border-b border-gray-800 min-w-[80px] sticky left-0 z-30 bg-[#15171c] border-r">
                  Código
                </th>
                <th 
                  className="px-4 py-3 border-b border-gray-800 min-w-[200px] cursor-pointer hover:bg-white/5 transition-colors"
                  onClick={() => handleSort('name')}
                >
                  Nome {getSortIcon('name')}
                </th>
                <th 
                  className="px-4 py-3 border-b border-gray-800 min-w-[200px] cursor-pointer hover:bg-white/5 transition-colors"
                  onClick={() => handleSort('company')}
                >
                  Empresa {getSortIcon('company')}
                </th>
                <th 
                  className="px-4 py-3 border-b border-gray-800 min-w-[150px] cursor-pointer hover:bg-white/5 transition-colors"
                  onClick={() => handleSort('region')}
                >
                  Região {getSortIcon('region')}
                </th>
                <th 
                  className="px-4 py-3 border-b border-gray-800 min-w-[150px] cursor-pointer hover:bg-white/5 transition-colors"
                  onClick={() => handleSort('city')}
                >
                  Cidade {getSortIcon('city')}
                </th>
                <th 
                  className="px-4 py-3 border-b border-gray-800 min-w-[200px] cursor-pointer hover:bg-white/5 transition-colors"
                  onClick={() => handleSort('email')}
                >
                  Email {getSortIcon('email')}
                </th>
                <th className="px-4 py-3 border-b border-gray-800 min-w-[100px] text-center">
                  Status
                </th>
                <th className="px-4 py-3 border-b border-gray-800 min-w-[150px] text-center bg-[#15171c] sticky right-0 z-30 border-l">
                  Ações
                </th>
              </tr>
            </thead>
            <tbody className="bg-[#1a1d24]">
              {data.data.map((representative) => (
                <tr key={representative.uuid} className="hover:bg-white/5 transition-colors border-b border-gray-800/20">
                  <td className="px-4 py-3 sticky left-0 z-10 bg-[#1a1d24] border-r border-gray-800/30 font-mono text-gray-300">
                    {representative.code}
                  </td>
                  <td className="px-4 py-3 font-medium text-gray-100">
                    {representative.name}
                  </td>
                  <td className="px-4 py-3 text-gray-300">
                    {representative.company}
                  </td>
                  <td className="px-4 py-3 text-gray-300">
                    {representative.region}
                  </td>
                  <td className="px-4 py-3 text-gray-300">
                    {representative.city}
                  </td>
                  <td className="px-4 py-3 text-gray-400 text-xs">
                    {representative.email}
                  </td>
                  <td className="px-4 py-3 text-center">
                    <span className={`inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium ${
                      representative.active
                        ? 'bg-emerald-500/10 text-emerald-400 border border-emerald-500/20'
                        : 'bg-rose-500/10 text-rose-400 border border-rose-500/20'
                    }`}>
                      {representative.active ? 'Ativo' : 'Inativo'}
                    </span>
                  </td>
                  <td className="px-4 py-3 text-center bg-[#1a1d24] sticky right-0 z-10 border-l border-gray-800/30">
                    <div className="flex items-center justify-center gap-2">
                      <button
                        onClick={() => onView(representative)}
                        className="p-1.5 hover:bg-blue-500/20 rounded-lg transition-colors text-blue-400 hover:text-blue-300"
                        title="Ver detalhes"
                      >
                        <Eye size={16} />
                      </button>
                      <button
                        onClick={() => onEdit(representative)}
                        className="p-1.5 hover:bg-amber-500/20 rounded-lg transition-colors text-amber-400 hover:text-amber-300"
                        title="Editar"
                      >
                        <Edit2 size={16} />
                      </button>
                      <button
                        onClick={() => onDelete(representative.uuid)}
                        className="p-1.5 hover:bg-rose-500/20 rounded-lg transition-colors text-rose-400 hover:text-rose-300"
                        title="Excluir"
                      >
                        <Trash2 size={16} />
                      </button>
                    </div>
                  </td>
                </tr>
              ))}
              {data.data.length === 0 && (
                <tr>
                  <td colSpan={8} className="px-4 py-12 text-center text-gray-500">
                    Nenhum representante encontrado
                  </td>
                </tr>
              )}
            </tbody>
          </table>
        </div>

        {/* Pagination */}
        {data.totalPages > 1 && (
          <div className="bg-[#15171c] px-4 py-3 border-t border-gray-800 flex items-center justify-between">
            <div className="text-sm text-gray-400">
              Mostrando {((data.page - 1) * data.pageSize) + 1} - {Math.min(data.page * data.pageSize, data.total)} de {data.total}
            </div>
            <div className="flex items-center gap-2">
              <select
                value={localLimit}
                onChange={(e) => handleLimitChange(Number(e.target.value))}
                className="px-3 py-1.5 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 text-sm focus:outline-none focus:ring-2 focus:ring-[#1e5144]"
              >
                <option value={10}>10 por página</option>
                <option value={20}>20 por página</option>
                <option value={50}>50 por página</option>
                <option value={100}>100 por página</option>
              </select>
              <button
                onClick={() => handlePageChange(Math.max(1, localPage - 1))}
                disabled={localPage === 1}
                className="px-3 py-1.5 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 text-sm hover:bg-white/5 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                Anterior
              </button>
              <span className="text-sm text-gray-400">
                Página {localPage} de {data.totalPages}
              </span>
              <button
                onClick={() => handlePageChange(Math.min(data.totalPages, localPage + 1))}
                disabled={localPage === data.totalPages}
                className="px-3 py-1.5 bg-[#1a1d24] border border-gray-700 rounded-lg text-gray-200 text-sm hover:bg-white/5 disabled:opacity-50 disabled:cursor-not-allowed transition-colors"
              >
                Próxima
              </button>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default RepresentativeTable;
