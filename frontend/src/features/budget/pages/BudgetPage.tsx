import React from 'react';
import BudgetView from '../ui/BudgetView';

const BudgetPage: React.FC = () => {
  const handleBack = () => {
    // TODO: Navigate back
    console.log('Navigate back');
  };

  return <BudgetView onBack={handleBack} />;
};

export default BudgetPage;
