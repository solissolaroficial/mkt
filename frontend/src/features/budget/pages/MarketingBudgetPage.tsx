import React, { useState } from 'react';
import MarketingBudgetView from '../ui/MarketingBudgetView';

const MarketingBudgetPage: React.FC = () => {
  const [selectedMonth, setSelectedMonth] = useState('NOV');
  const handleBack = () => {
    // TODO: Navigate back
    console.log('Navigate back');
  };

  return <MarketingBudgetView selectedMonth={selectedMonth} onBack={handleBack} />;
};

export default MarketingBudgetPage;
