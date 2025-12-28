import React from 'react';
import SocialBenchmarkingView from '../ui/SocialBenchmarkingView';
import { SOCIAL_MEDIA_RANKING } from '@/shared/utils/legacy.constants';

const SocialPage: React.FC = () => {
  return <SocialBenchmarkingView data={SOCIAL_MEDIA_RANKING} />;
};

export default SocialPage;
