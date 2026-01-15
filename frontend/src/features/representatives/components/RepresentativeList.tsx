import React, { useState } from 'react';
import { useRepresentativeMutations } from '../hooks/useRepresentativeMutations';
import { useRepStats } from '../hooks/useRepStats';
import RepresentativeTable from './RepresentativeTable';
import RepresentativeForm from './RepresentativeForm';
import RepresentativeStats from './RepresentativeStats';
import type { Representative, ListRepresentativesRequest, RepresentativeStats as RepresentativeStatsType } from '../types';

interface RepresentativeListProps {
  data: {
    data: Representative[];
    total: number;
    page: number;
    pageSize: number;
    totalPages: number;
  };
  isLoading: boolean;
  error: unknown;
  onRefetch: () => void;
  onPageChange: (page: number) => void;
  onLimitChange: (limit: number) => void;
  onSearch: (term: string) => void;
  onSortChange: (sortBy: string, sortOrder: 'asc' | 'desc') => void;
  currentFilters?: ListRepresentativesRequest;
}

const RepresentativeList: React.FC<RepresentativeListProps> = ({
  data,
  isLoading,
  error,
  onRefetch,
  onPageChange,
  onLimitChange,
  onSearch,
  onSortChange,
  currentFilters,
}) => {
  const [selectedRepresentative, setSelectedRepresentative] = useState<Representative | null>(null);
  const [showForm, setShowForm] = useState(false);
  const [showStats, setShowStats] = useState(false);
  const [statsRepresentative, setStatsRepresentative] = useState<Representative | null>(null);

  const { createItem, updateItem, deleteItem, isCreating, isUpdating, isDeleting } = useRepresentativeMutations();
  const { data: statsData, isLoading: statsLoading, error: statsError, refetch: refetchStats } = useRepStats(statsRepresentative?.uuid || '');

  const handleCreate = () => {
    setSelectedRepresentative(null);
    setShowForm(true);
  };

  const handleEdit = (representative: Representative) => {
    setSelectedRepresentative(representative);
    setShowForm(true);
  };

  const handleViewStats = (representative: Representative) => {
    setStatsRepresentative(representative);
    setShowStats(true);
  };

  const handleDelete = (uuid: string) => {
    if (confirm('Tem certeza que deseja excluir este representante?')) {
      deleteItem(uuid);
    }
  };

  const handleFormSubmit = (data: any) => {
    if (selectedRepresentative) {
      updateItem(
        { uuid: selectedRepresentative.uuid, data },
        {
          onSuccess: () => {
            setShowForm(false);
            setSelectedRepresentative(null);
            onRefetch();
          },
        }
      );
    } else {
      createItem(data, {
        onSuccess: () => {
          setShowForm(false);
          onRefetch();
        },
      });
    }
  };

  return (
    <>
      {/* Table View */}
      <RepresentativeTable
        data={data}
        isLoading={isLoading}
        error={error}
        onRefetch={onRefetch}
        onCreate={handleCreate}
        onEdit={handleEdit}
        onView={handleViewStats}
        onDelete={handleDelete}
        onPageChange={onPageChange}
        onLimitChange={onLimitChange}
        onSearch={onSearch}
        onSortChange={onSortChange}
        currentFilters={currentFilters}
      />

      {/* Create/Edit Form Modal */}
      {showForm && (
        <RepresentativeForm
          representative={selectedRepresentative || undefined}
          onSubmit={handleFormSubmit}
          onCancel={() => {
            setShowForm(false);
            setSelectedRepresentative(null);
          }}
          isLoading={isCreating || isUpdating}
        />
      )}

      {/* Stats Modal */}
      {showStats && statsRepresentative && (
        <RepresentativeStats
          stats={statsData || {
            uuid: statsRepresentative.uuid,
            onlineActionCount: 0,
            offlineActionCount: 0,
            offlineActionValue: 0,
            showroomItemCount: 0,
            repMarketingCount: 0,
            totalActions: 0,
            onlineCount: 0,
            offlineCount: 0,
            offlineValue: 0,
          }}
          representative={statsRepresentative}
          isLoading={statsLoading}
          error={statsError}
          onRefetch={refetchStats}
          onBack={() => {
            setShowStats(false);
            setStatsRepresentative(null);
          }}
        />
      )}
    </>
  );
};

export default RepresentativeList;
