import React from 'react';
import SocialBenchmarkingView from '../ui/SocialBenchmarkingView';
import { useSocialBenchmarkings } from '../hooks/useSocial';
import { Loader2 } from 'lucide-react';

const SocialPage: React.FC = () => {
  const { data, isLoading, error } = useSocialBenchmarkings();

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

  return <SocialBenchmarkingView data={data?.benchmarkings || []} />;
};

export default SocialPage;
