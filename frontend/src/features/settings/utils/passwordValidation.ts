export interface PasswordValidationResult {
  isValid: boolean;
  errors: string[];
  strength: 'weak' | 'medium' | 'strong';
}

export function validatePassword(password: string): PasswordValidationResult {
  const errors: string[] = [];
  
  // Verificar comprimento mínimo
  if (password.length < 8) {
    errors.push('A senha deve ter no mínimo 8 caracteres');
  }
  
  // Verificar se tem pelo menos 3 dos 4 tipos de caracteres
  const hasUpperCase = /[A-Z]/.test(password);
  const hasLowerCase = /[a-z]/.test(password);
  const hasNumbers = /[0-9]/.test(password);
  const hasSpecial = /[!@#$%^&*(),.?":{}|<>]/.test(password);
  
  const characterTypeCount = [hasUpperCase, hasLowerCase, hasNumbers, hasSpecial].filter(Boolean).length;
  
  if (characterTypeCount < 3) {
    errors.push('A senha deve conter pelo menos 3 dos seguintes tipos: letras maiúsculas, letras minúsculas, números ou caracteres especiais');
  }
  
  // Calcular força da senha
  let strength: 'weak' | 'medium' | 'strong' = 'weak';
  if (password.length >= 8 && characterTypeCount >= 3) {
    if (password.length >= 12 && characterTypeCount === 4) {
      strength = 'strong';
    } else {
      strength = 'medium';
    }
  }
  
  return {
    isValid: errors.length === 0,
    errors,
    strength,
  };
}

export function getPasswordStrengthColor(strength: 'weak' | 'medium' | 'strong'): string {
  switch (strength) {
    case 'weak':
      return '#ef4444'; // red
    case 'medium':
      return '#f59e0b'; // yellow
    case 'strong':
      return '#10b981'; // green
  }
}

export function getPasswordStrengthLabel(strength: 'weak' | 'medium' | 'strong'): string {
  switch (strength) {
    case 'weak':
      return 'Fraca';
    case 'medium':
      return 'Média';
    case 'strong':
      return 'Forte';
  }
}
