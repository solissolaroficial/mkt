import React, { useState } from 'react';
import { useQueryClient } from '@tanstack/react-query';
import SocialBenchmarkingView from '../ui/SocialBenchmarkingView';
import SocialPostsView from '../ui/SocialPostsView';
import SocialDailyAggregationsView from '../ui/SocialDailyAggregationsView';
import BrandsView from '../ui/BrandsView';
import { useSocialBenchmarkings, useSocialPosts, useSocialDailyAggregations } from '../hooks/useSocial';
import { Loader2, BarChart2, Image, Calendar, Tag } from 'lucide-react';
 
type TabType = 'brands' | 'benchmarking' | 'posts' | 'aggregations';

const SocialPage: React.FC = () => {
  const [activeTab, setActiveTab] = useState<TabType>('brands');
  const queryClient = useQueryClient();
  
  const { data: benchmarkingData, isLoading: isLoadingBenchmarking, error: benchmarkingError } = useSocialBenchmarkings();
  const { data: postsData, isLoading: isLoadingPosts, error: postsError } = useSocialPosts();
  const { data: aggregationsData, isLoading: isLoadingAggregations, error: aggregationsError } = useSocialDailyAggregations();
  
  const isLoading = isLoadingBenchmarking || isLoadingPosts || isLoadingAggregations;
  const error = benchmarkingError || postsError || aggregationsError;
 
  if (isLoading) {
    return (
      <div className="flex items-center justify-center h-full">
        <Loader2 className="animate-spin text-gray-400" size={48} />
      </div>
    );
  }
 
  if (error) {
    return (
      <div className="flex items-center justify-center h-full">
        <div className="text-center">
          <p className="text-red-400 text-lg mb-2">Erro ao carregar dados</p>
          <p className="text-gray-500 text-sm">Tente novamente mais tarde</p>
        </div>
      </div>
    );
  }

  const tabs: { id: TabType; label: string; icon: React.ReactNode }[] = [
    { id: 'brands', label: 'Marcas', icon: <Tag size={16} /> },
    { id: 'benchmarking', label: 'Benchmarking', icon: <BarChart2 size={16} /> },
    { id: 'posts', label: 'Posts', icon: <Image size={16} /> },
    { id: 'aggregations', label: 'Agregações', icon: <Calendar size={16} /> },
  ];

  return (
    <div className="h-full flex flex-col">
      {/* Tabs */}
      <div className="flex gap-2 mb-6 border-b border-gray-800">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            onClick={() => setActiveTab(tab.id)}
            className={`
              flex items-center gap-2 px-4 py-3 text-sm font-medium transition-colors border-b-2 -mb-px
              ${activeTab === tab.id
                ? 'text-white border-[#1e5144] bg-[#1e5144]/10'
                : 'text-gray-500 border-transparent hover:text-gray-400 hover:bg-[#1a1d24]'
              }
            `}
          >
            {tab.icon}
            {tab.label}
          </button>
        ))}
      </div>

      {/* Content */}
      <div className="flex-grow">
        {activeTab === 'brands' && (
          <BrandsView
            onRefresh={() => {
              // Trigger refetch by invalidating queries
              queryClient.invalidateQueries({ queryKey: ['brands'] });
            }}
          />
        )}
        {activeTab === 'benchmarking' && (
          <SocialBenchmarkingView data={benchmarkingData?.benchmarkings || []} />
        )}
        {activeTab === 'posts' && (
          <SocialPostsView
            data={postsData?.data?.posts || []}
            onRefresh={() => {
              // Trigger refetch by invalidating queries
              queryClient.invalidateQueries({ queryKey: ['social-posts'] });
            }}
          />
        )}
        {activeTab === 'aggregations' && (
          <SocialDailyAggregationsView
            data={aggregationsData?.data?.aggregations || []}
            onRefresh={() => {
              // Trigger refetch by invalidating queries
              queryClient.invalidateQueries({ queryKey: ['social-daily-aggregations'] });
            }}
          />
        )}
      </div>
    </div>
  );
};
 
export default SocialPage;
