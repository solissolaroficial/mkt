import React, { useState } from 'react';
import { useRepTableData } from '../hooks/useRepTableData';
import RepresentativeList from '../components/RepresentativeList';
import { MonthlyGoalsTable } from '../components/MonthlyGoalsTable';
import RepProfileModal from '../ui/RepProfileModal';
import { representativesService } from '../services/representativesService';
import type { ListRepresentativesRequest, RepresentativeProfile } from '../types';

const RepresentativesPage: React.FC = () => {
  const [viewMode, setViewMode] = useState<'list' | 'table'>('list');
  const [filters, setFilters] = useState<ListRepresentativesRequest>({
    page: 1,
    limit: 10,
    sortBy: 'name',
    sortOrder: 'asc',
  });
  const [selectedProfile, setSelectedProfile] = useState<RepresentativeProfile | null>(null);
  const [isProfileModalOpen, setIsProfileModalOpen] = useState(false);

  const { data, isLoading, error, refetch } = useRepTableData(filters);

  const handlePageChange = (newPage: number) => {
    setFilters(prev => ({ ...prev, page: newPage }));
  };

  const handleLimitChange = (newLimit: number) => {
    setFilters(prev => ({ ...prev, limit: newLimit, page: 1 }));
  };

  const handleSearch = (term: string) => {
    setFilters(prev => ({
      ...prev,
      name: term || undefined,
      page: 1,
    }));
  };

  const handleSortChange = (sortBy: string, sortOrder: 'asc' | 'desc') => {
    setFilters(prev => ({ ...prev, sortBy, sortOrder }));
  };

  const handleRepClick = async (repName: string) => {
    try {
      const profile = await representativesService.getRepProfile(repName);
      setSelectedProfile(profile);
      setIsProfileModalOpen(true);
    } catch (error) {
      console.error('Error fetching representative profile:', error);
    }
  };

  const handleCloseProfileModal = () => {
    setIsProfileModalOpen(false);
    setSelectedProfile(null);
  };

  return (
    <div className="space-y-6">
      {/* View Toggle */}
      <div className="flex items-center justify-between">
        <h1 className="text-2xl font-bold text-gray-900">Representantes</h1>
        <div className="flex space-x-2">
          <button
            onClick={() => setViewMode('list')}
            className={`px-4 py-2 rounded-md text-sm font-medium ${
              viewMode === 'list'
                ? 'bg-blue-600 text-white'
                : 'bg-white text-gray-700 border border-gray-300 hover:bg-gray-50'
            }`}
          >
            Lista
          </button>
          <button
            onClick={() => setViewMode('table')}
            className={`px-4 py-2 rounded-md text-sm font-medium ${
              viewMode === 'table'
                ? 'bg-blue-600 text-white'
                : 'bg-white text-gray-700 border border-gray-300 hover:bg-gray-50'
            }`}
          >
            Tabela de Metas
          </button>
        </div>
      </div>

      {/* List View */}
      {viewMode === 'list' && (
        <RepresentativeList
          data={{
            data: data?.data || [],
            total: data?.total || 0,
            page: data?.page || 1,
            pageSize: data?.pageSize || 10,
            totalPages: data?.totalPages || 1,
          }}
          isLoading={isLoading}
          error={error}
          onRefetch={refetch}
          onPageChange={handlePageChange}
          onLimitChange={handleLimitChange}
          onSearch={handleSearch}
          onSortChange={handleSortChange}
          currentFilters={filters}
        />
      )}

      {/* Table View - Monthly Goals */}
      {viewMode === 'table' && (
        <div className="space-y-4">
          <div className="flex items-center space-x-4">
            <label className="text-sm font-medium text-gray-700">Ano:</label>
            <select
              className="border border-gray-300 rounded-md px-3 py-2 text-sm"
              defaultValue={new Date().getFullYear()}
            >
              <option value={2024}>2024</option>
              <option value={2025}>2025</option>
              <option value={2026}>2026</option>
            </select>
          </div>
          <MonthlyGoalsTable year={new Date().getFullYear()} onRepClick={handleRepClick} />
        </div>
      )}

      {/* Profile Modal */}
      {isProfileModalOpen && selectedProfile && (
        <RepProfileModal
          isOpen={isProfileModalOpen}
          onClose={handleCloseProfileModal}
          profile={selectedProfile}
        />
      )}
    </div>
  );
};

export default RepresentativesPage;
