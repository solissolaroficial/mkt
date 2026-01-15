import React, { useState } from 'react';
import { useBudgetItems, useBudgetYears } from '../hooks/useBudget';
import BudgetView from '../ui/BudgetView';

const BudgetPage: React.FC = () => {
  // Chamar hooks para carregar dados
  const [selectedYear, setSelectedYear] = useState<number | undefined>(undefined);
  const { data: years = [] } = useBudgetYears();
  const { data: budgetItems = [], isLoading, error, refetch } = useBudgetItems({
    year: selectedYear,
  });

  const handleBack = () => {
    // TODO: Navigate back
    console.log('Navigate back');
  };

  // Passar dados para BudgetView
  return (
    <BudgetView
      onBack={handleBack}
      dataItems={budgetItems}
      isLoading={isLoading}
      error={error}
      onRefetch={refetch}
      years={years}
      selectedYear={selectedYear}
      onYearChange={setSelectedYear}
    />
  );
};

export default BudgetPage;
