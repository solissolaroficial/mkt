import axios from 'axios';

const API_URL = import.meta.env.VITE_API_URL || 'http://localhost:3000';

export const apiClient = axios.create({
  baseURL: API_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Request interceptor - adiciona token JWT automaticamente
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

// Response interceptor - trata erro 401 (não autorizado)
apiClient.interceptors.response.use(
  (response) => {
    console.log('[AXIOS INTERCEPTOR] Response:', response);
    console.log('[AXIOS INTERCEPTOR] Response.data:', response.data);
    console.log('[AXIOS INTERCEPTOR] Response.config.url:', response.config.url);
    return response;
  },
  (error) => {
    console.error('[AXIOS INTERCEPTOR] Error:', error);
    if (error.response?.status === 401) {
      // Remove token inválido e redireciona para login
      localStorage.removeItem('auth_token');
      window.location.href = '/login';
    }
    return Promise.reject(error);
  }
);