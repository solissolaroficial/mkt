import React, { useState } from 'react';
import type { Flow } from '@/shared/types';
import { useFlows, useFlowMutations } from '../hooks';
import { FlowModal } from './FlowModal';
import { Plus, Edit2, Trash2, GripVertical } from 'lucide-react';
import {
  DndContext,
  closestCenter,
  KeyboardSensor,
  PointerSensor,
  useSensor,
  useSensors,
  DragEndEvent,
} from '@dnd-kit/core';
import {
  arrayMove,
  SortableContext,
  sortableKeyboardCoordinates,
  useSortable,
  verticalListSortingStrategy,
} from '@dnd-kit/sortable';
import { CSS } from '@dnd-kit/utilities';

interface FlowSidebarProps {
  selectedFlowId?: string;
  onSelectFlow: (flowId: string | null) => void;
}

// Componente SortableFlowItem para drag and drop
const SortableFlowItem = ({
  flow,
  selectedFlowId,
  onEdit,
  onDelete,
  onSelectFlow
}: {
  flow: Flow;
  selectedFlowId?: string;
  onEdit: (flow: Flow) => void;
  onDelete: (flowId: string, flowName: string, e: React.MouseEvent) => void;
  onSelectFlow?: (flowUuid: string) => void;
}) => {
  const {
    attributes,
    listeners,
    setNodeRef,
    transform,
    transition,
    isDragging,
  } = useSortable({ id: flow.uuid });

  const style = {
    transform: CSS.Transform.toString(transform),
    transition,
    opacity: isDragging ? 0.5 : 1,
  };

  const handleEditClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onEdit(flow);
  };

  const handleDeleteClick = (e: React.MouseEvent) => {
    e.stopPropagation();
    onDelete(flow.uuid, flow.name, e);
  };

  return (
    <div
      id={flow.uuid}
      ref={setNodeRef}
      style={style}
      onClick={() => onSelectFlow && onSelectFlow(flow.uuid)}
      className={`w-full flex items-center gap-2 px-3 py-2.5 rounded-lg transition-all group ${selectedFlowId === flow.uuid
          ? 'bg-[#20232b] border border-gray-700'
          : 'bg-transparent hover:bg-[#1a1d24] border border-transparent'
        }`}
    >
      <div
        {...attributes}
        {...listeners}
        className="cursor-grab active:cursor-grabbing"
      >
        <GripVertical size={14} className="text-gray-600" />
      </div>

      <div
        className="w-3 h-3 rounded-full flex-shrink-0"
        style={{ backgroundColor: flow.color }}
      />

      <div className="flex-1 min-w-0">
        <p className={`text-sm font-medium truncate ${selectedFlowId === flow.uuid ? 'text-white' : 'text-gray-300'
          }`}>
          {flow.name}
        </p>
        {flow.description && (
          <p className="text-xs text-gray-500 truncate">
            {flow.description}
          </p>
        )}
      </div>

      <div className="flex items-center gap-1 opacity-0 group-hover:opacity-100 transition-opacity">
        <button
          onClick={handleEditClick}
          className="p-1.5 hover:bg-white/10 rounded transition-colors"
        >
          <Edit2 size={14} className="text-gray-400" />
        </button>
        <button
          onClick={handleDeleteClick}
          className="p-1.5 hover:bg-white/10 rounded transition-colors"
        >
          <Trash2 size={14} className="text-gray-400" />
        </button>
      </div>
    </div>
  );
};

export const FlowSidebar: React.FC<FlowSidebarProps> = ({
  selectedFlowId,
  onSelectFlow,
}) => {
  const { data: flowsData, isLoading } = useFlows();
  const { createFlow, updateFlow, deleteFlow, reorderFlows, isCreating, isUpdating, isDeleting } = useFlowMutations();

  const [isModalOpen, setIsModalOpen] = useState(false);
  const [editingFlow, setEditingFlow] = useState<Flow | null>(null);
  const [deleteConfirm, setDeleteConfirm] = useState<{ flowId: string; flowName: string } | null>(null);

  const flows = flowsData?.data || [];

  // Debug: verificar IDs duplicados
  console.log('FlowSidebar: Flows data:', flows);
  console.log('FlowSidebar: Flow IDs:', flows.map(f => f.uuid));
  const uniqueIds = new Set(flows.map(f => f.uuid));
  console.log('FlowSidebar: Unique IDs count:', uniqueIds.size, 'Total flows:', flows.length);
  if (uniqueIds.size !== flows.length) {
    console.error('FlowSidebar: DUPLICATE IDS FOUND!');
  }

  // Configuração de sensors para drag and drop
  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  const handleCreateFlow = async (data: { name: string; description?: string; color?: string }) => {
    await createFlow(data);
  };

  const handleUpdateFlow = async (data: { name: string; description?: string; color?: string }) => {
    if (editingFlow) {
      await updateFlow({ id: editingFlow.uuid, data });
    }
  };

  const handleEditFlow = (flow: Flow) => {
    console.log('FlowSidebar: handleEditFlow called with flow:', flow);
    setEditingFlow(flow);
    setIsModalOpen(true);
  };

  const handleDeleteFlow = (flowId: string, flowName: string, e: React.MouseEvent) => {
    e.stopPropagation();
    setDeleteConfirm({ flowId, flowName });
  };

  const handleSelectFlow = (flowUuid: string) => {
    onSelectFlow(flowUuid);
  };

  // Handler para drag end
  const handleDragEnd = (event: DragEndEvent) => {
    const { active, over } = event;

    if (over && active.id !== over.id) {
      const oldIndex = flows.findIndex((flow) => flow.uuid === active.id);
      const newIndex = flows.findIndex((flow) => flow.uuid === over.id);
      const newFlows = arrayMove(flows, oldIndex, newIndex);

      console.log('FlowSidebar: Reordering flows from index', oldIndex, 'to', newIndex);
      console.log('FlowSidebar: New order:', newFlows.map((f: Flow) => f.uuid));

      // Chamar a mutation de reordenar com os IDs na nova ordem
      reorderFlows(newFlows.map((f: Flow) => f.uuid));
    }
  };

  if (isLoading) {
    return (
      <div className="w-64 bg-[#0f1115] border-r border-gray-800 flex flex-col">
        <div className="p-4 border-b border-gray-800">
          <div className="h-6 bg-gray-800 rounded animate-pulse w-24" />
        </div>
        <div className="flex-1 overflow-y-auto p-4 space-y-3">
          {[1, 2, 3].map((i) => (
            <div key={i} className="h-12 bg-gray-800/50 rounded animate-pulse" />
          ))}
        </div>
      </div>
    );
  }

  return (
    <>
      <div className="w-64 bg-[#0f1115] border-r border-gray-800 flex flex-col">
        <div className="p-4 border-b border-gray-800">
          <h2 className="text-sm font-bold text-gray-300 uppercase tracking-wider mb-3">
            Fluxos
          </h2>
          <button
            onClick={() => {
              setEditingFlow(null);
              setIsModalOpen(true);
            }}
            disabled={isCreating}
            className="w-full flex items-center justify-center gap-2 px-3 py-2.5 bg-[#20232b] text-gray-300 text-sm font-medium rounded-lg hover:bg-[#2a2e37] disabled:opacity-50 transition-colors"
          >
            <Plus size={16} />
            Novo Fluxo
          </button>
        </div>

        <div className="flex-1 overflow-y-auto p-3 space-y-2">
          {flows.length === 0 ? (
            <div className="text-center py-8 text-gray-500">
              <p className="text-sm">Nenhum fluxo criado ainda</p>
            </div>
          ) : (
            <DndContext
              sensors={sensors}
              collisionDetection={closestCenter}
              onDragEnd={handleDragEnd}
            >
              <SortableContext items={flows.map(flow => flow.uuid)} strategy={verticalListSortingStrategy}>
                {flows.map((flow) => (
                  <SortableFlowItem
                    key={flow.uuid}
                    flow={flow}
                    selectedFlowId={selectedFlowId}
                    onEdit={handleEditFlow}
                    onDelete={handleDeleteFlow}
                    onSelectFlow={handleSelectFlow}
                  />
                ))}
              </SortableContext>
            </DndContext>
          )}
        </div>

        <div className="p-4 border-t border-gray-800 text-xs text-gray-600">
          <p>{flows.length} fluxo(s)</p>
        </div>
      </div>

      <FlowModal
        isOpen={isModalOpen}
        onClose={() => {
          setIsModalOpen(false);
          setEditingFlow(null);
        }}
        onSave={editingFlow ? handleUpdateFlow : handleCreateFlow}
        editingFlow={editingFlow}
      />

      {/* --- DELETE CONFIRMATION MODAL --- */}
      {deleteConfirm && (
        <div className="fixed inset-0 bg-black/80 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
          <div className="bg-[#1a1d24] w-full max-w-md rounded-2xl shadow-2xl border border-gray-700 overflow-hidden">
            <div className="p-6 border-b border-gray-800">
              <h3 className="text-lg font-bold text-gray-100">Confirmar Remoção</h3>
              <p className="text-sm text-gray-400 mt-2">
                Tem certeza que deseja excluir o fluxo <span className="font-bold text-white">{deleteConfirm.flowName}</span>?
              </p>
              <p className="text-xs text-rose-400 mt-1">
                Todas as tarefas deste fluxo serão arquivadas.
              </p>
            </div>
            <div className="p-6 flex justify-end gap-3 bg-[#20232b]">
              <button
                onClick={() => setDeleteConfirm(null)}
                className="px-4 py-2 text-gray-400 hover:bg-gray-800 rounded-lg text-sm font-medium"
              >
                Cancelar
              </button>
              <button
                onClick={() => {
                  if (deleteConfirm.flowId) {
                    deleteFlow(deleteConfirm.flowId);
                    setDeleteConfirm(null);
                  }
                }}
                disabled={isDeleting}
                className="px-4 py-2 bg-rose-600 text-white hover:bg-rose-700 rounded-lg text-sm font-medium shadow-lg disabled:opacity-50"
              >
                {isDeleting ? 'Removendo...' : 'Remover'}
              </button>
            </div>
          </div>
        </div>
      )}
    </>
  );
};
