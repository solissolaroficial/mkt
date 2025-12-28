import React from 'react';
import RepTable from '../ui/RepTable';
import { REPS_TRAINING_DATA } from '@/shared/utils/legacy.constants';

const RepresentativesPage: React.FC = () => {
  return <RepTable data={REPS_TRAINING_DATA} />;
};

export default RepresentativesPage;
