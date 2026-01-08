import React, { useState, useMemo } from 'react';
import PostCalendarView from '../ui/PostCalendarView';
import { useCalendarPosts } from '../hooks/useCalendar';
import { useCalendarMutations } from '../hooks/useCalendarMutations';
import { ErrorDisplay } from '@/shared/components/ErrorDisplay';
import type { PostCategory, PostType, CalendarPost, PublicUser } from '@/shared/types/legacy.types';
import { useUsers } from '@/features/users/hooks';
import type { AppUser } from '@/shared/types/user.types';

const CalendarPage: React.FC = () => {
  // Filtros gerenciados no nível da página
  const [selectedCategory, setSelectedCategory] = useState<PostCategory | 'all'>('all');
  const [selectedType, setSelectedType] = useState<PostType | 'all'>('all');
  const [selectedAssignee, setSelectedAssignee] = useState<string>('all');
  
  // Buscar posts do backend usando React Query com filtros
  const { data: postsData, isLoading, error, refetch } = useCalendarPosts({
    category: selectedCategory === 'all' ? undefined : selectedCategory,
    type: selectedType === 'all' ? undefined : selectedType,
    assignee_id: selectedAssignee === 'all' ? undefined : selectedAssignee,
  });
  
  // Buscar usuários do sistema
  const { data: appUsers } = useUsers();
  
  // Obter mutations do calendário
  const mutations = useCalendarMutations();

  // Popular o objeto assignee usando os dados do useUsers hook
  const posts = useMemo(() => {
    const rawPosts = postsData?.data || [];
    
    console.log('[CALENDAR PAGE] rawPosts:', rawPosts);
    console.log('[CALENDAR PAGE] rawPosts length:', rawPosts?.length);
    console.log('[CALENDAR PAGE] First post date:', rawPosts?.[0]?.date);
    
    // Criar um mapa de usuários por UUID para busca rápida
    const usersMap = new Map<string, AppUser>();
    if (appUsers) {
      appUsers.forEach(user => {
        usersMap.set(user.id, user);
      });
    }
    
    // Mapear posts para popular o objeto assignee
    const mappedPosts = rawPosts.map(post => {
      console.log('[CALENDAR PAGE] Processing post:', post.id, 'date:', post.date);
      
      // Se o post já tem um objeto assignee, retornar como está
      if (post.assignee) {
        return post;
      }
      
      // Se não tem assignee mas tem assignee_id, buscar no mapa de usuários
      if (post.assignee_id && usersMap.has(post.assignee_id)) {
        const user = usersMap.get(post.assignee_id)!;
        const assignee: PublicUser = {
          id: user.id,
          name: user.name,
          email: user.email
        };
        
        return {
          ...post,
          assignee
        };
      }
      
      // Se não tem assignee_id ou não encontrou o usuário, retornar post original
      return post;
    });
    
    console.log('[CALENDAR PAGE] mappedPosts:', mappedPosts);
    console.log('[CALENDAR PAGE] mappedPosts length:', mappedPosts?.length);
    console.log('[CALENDAR PAGE] First mapped post date:', mappedPosts?.[0]?.date);
    
    return mappedPosts;
  }, [postsData?.data, appUsers]);

  // Estado de loading simples (texto)
  if (isLoading) {
    return <div>Carregando calendário...</div>;
  }

  // Estado de erro
  if (error) {
    return <ErrorDisplay error={error} onRetry={refetch} />;
  }

  return (
    <PostCalendarView
      posts={posts}
      pagination={postsData}
      mutations={mutations}
      selectedCategory={selectedCategory}
      setSelectedCategory={setSelectedCategory}
      selectedType={selectedType}
      setSelectedType={setSelectedType}
      selectedAssignee={selectedAssignee}
      setSelectedAssignee={setSelectedAssignee}
    />
  );
};

export default CalendarPage;
