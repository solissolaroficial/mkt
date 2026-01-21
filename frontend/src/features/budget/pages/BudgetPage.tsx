import React from 'react';
import { useBudgetItems, useBudgetYears } from '../hooks/useBudget';
import { useUIStore } from '@/shared/store/uiStore';
import BudgetView from '../ui/BudgetView';

const BudgetPage: React.FC = () => {
  // Chamar hooks para carregar dados
  const { selectedYear, setSelectedYear } = useUIStore();
  const { data: years = [] } = useBudgetYears();
  const { data: budgetItems = [], isLoading, error, refetch } = useBudgetItems({
    year: selectedYear ? parseInt(selectedYear) : undefined,
  });

  const handleBack = () => {
    // TODO: Navigate back
    console.log('Navigate back');
  };

  // Converter tipos para BudgetView
  const selectedYearNumber = selectedYear ? parseInt(selectedYear) : undefined;
  const handleYearChange = (year: number | undefined) => {
    setSelectedYear(year ? year.toString() : '');
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
      selectedYear={selectedYearNumber}
      onYearChange={handleYearChange}
    />
  );
};

export default BudgetPage;
