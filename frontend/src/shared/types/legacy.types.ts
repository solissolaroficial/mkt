
export interface BreakdownItem {
  label: string;
  value: number;
  subItems?: BreakdownItem[];
}

export interface MonthlyData {
  year?: number; // Year (e.g., 2024, 2025)
  month: string;
  realized: number | null;
  meta: number | null;
  breakdown?: BreakdownItem[]; // Added for detailed view (e.g. Channel breakdown)
  logs?: KpiLog[]; // Logs are stored at MonthlyData level, not KPI level
}

export interface KpiLog {
  id: string;
  date: string;       // ISO string
  timestamp: string;  // Hora formatada
  user: string;
  month: string;
  oldValue: number | null;
  newValue: number;
  action: 'update' | 'create';
  context?: string;
}

export interface KpiCategory {
  id: string;
  title: string;
  slug: string;
  data: MonthlyData[];
  color: string;
  unit?: string; // 'currency' | 'percent' | 'number'
  logs?: KpiLog[];
}

export interface RepData {
  name: string;
  [month: string]: number | string; // Month keys with values
}

export interface RepTableData {
  title: string;
  columns: string[];
  rows: {
    month: string;
    target: number;
    values: Record<string, number>;
  }[];
}

export interface ChannelPerformance {
  channel: string;
  visits: number;
  leads: number;
  conversion: number; // percentage
  investment: number;
  cpl: number;
  roas: number;
}

export interface MarketingChannelData {
  [month: string]: ChannelPerformance[];
}

export type TaskStatus = 'pending' | 'in_progress' | 'completed' | 'cancelled';
export type TaskPriority = 'low' | 'medium' | 'high' | 'urgent';
export type TaskCategory = 'commercial' | 'marketing' | 'training' | 'administrative' | 'support';
export type TaskFlow = 'backlog' | 'to_do' | 'in_progress' | 'review' | 'done';
export type NotificationType = 'mention' | 'deadline' | 'system';
export type AllowedUser = TaskFlow;

export interface Subtask {
  id: string;
  task_id: string;
  title: string;
  description: string;
  priority: TaskPriority;
  status: TaskStatus;
  due_date?: string;
  assignee_id?: string;
  created_at: string;
  updated_at: string;
}

export interface Comment {
  id: string;
  task_id: string;
  user_id: string;
  text: string;
  timestamp: string;
}

export interface Notification {
  id: string;
  user_id: string;
  task_id?: string;
  notification_type: NotificationType;
  title: string;
  message: string;
  read: boolean;
  archived: boolean;
  timestamp: string;
}
export interface Task {
  id: string;
  title: string;
  description: string;
  category: TaskCategory;
  priority: TaskPriority;
  status: TaskStatus;
  flow: TaskFlow;
  due_date?: string;
  assignee_id?: string;
  flows: TaskFlow[];
  archived: boolean;
  subtasks?: Subtask[];
  comments?: Comment[];
  notifications?: Notification[];
  created_at: string;
  updated_at: string;
}

export interface PdvPost {
  id: string;
  representative_uuid: string;
  pdv_name: string;
  post_date: string; // YYYY-MM-DD
  month: string; // JAN, FEV, MAR, ABR, MAI, JUN, JUL, AGO, SET, OUT, NOV, DEZ
  platform: 'instagram' | 'facebook' | 'linkedin' | 'youtube' | 'tiktok';
  link?: string; // Opcional, max 500 caracteres
  proof_url?: string; // Opcional, max 500 caracteres
  status: 'pending' | 'approved' | 'rejected' | 'published' | 'cancelled';
  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601
}

export interface RecurrentPdv {
    id: string;
    name: string;
    representative_uuid: string;
    city?: string;
    followers?: number;
    instagram_profile?: string;
    created_at: string; // ISO 8601
    updated_at: string; // ISO 8601
}

// ============================================
// PDV Post Types (Backend Integration)
// ============================================

// Interface para resposta paginada de posts PDV
export interface PdvPostListResponse {
  posts: PdvPost[];
  meta: {
    total: number;
    page: number;
    limit: number;
    total_pages: number;
  };
}

// Interface para resposta paginada de PDVs recorrentes
export interface RecurrentPdvListResponse {
  pdvs: RecurrentPdv[];
  meta: {
    total: number;
    page: number;
    limit: number;
    total_pages: number;
  };
}

// Interface para filtros de posts PDV
export interface PdvPostFilters {
  representative_uuid?: string;
  month?: string;
  platform?: string;
  status?: string;
  start_date?: string;
  end_date?: string;
  page?: number;
  limit?: number;
  sort_by?: string;
  sort_order?: string;
}

// Interface para filtros de PDVs recorrentes
export interface RecurrentPdvFilters {
  representative_uuid?: string;
  city?: string;
  page?: number;
  limit?: number;
  sort_by?: string;
  sort_order?: string;
}

// Interface para request de criação de post PDV
export interface CreatePdvPostRequest {
  representative_uuid: string;
  pdv_name: string;
  post_date: string;
  platform: string;
  link?: string;
  proof_url?: string;
}

// Interface para request de atualização de post PDV
export interface UpdatePdvPostRequest {
  representative_uuid?: string;
  pdv_name?: string;
  post_date?: string;
  platform?: string;
  link?: string;
  proof_url?: string;
}

// Interface para request de criação de PDV recorrente
export interface CreateRecurrentPdvRequest {
  name: string;
  representative_uuid: string;
  city?: string;
}

// Interface para request de atualização de PDV recorrente
export interface UpdateRecurrentPdvRequest {
  name?: string;
  representative_uuid?: string;
  city?: string;
  followers?: number;
  instagram_profile?: string;
}

export interface ShowroomItem {
    id: string;
    pdv: string;
    city: string;
    contact: string;
    representative_uuid: string;
    notDelivered: boolean; // Coluna "Não entregue"
    deliveryForecast: string; // "Previsão de entrega"
    delivered: boolean; // Coluna "Já entregue"
    workshopDate: string; // "Data do Workshop"
}

export interface OfflineAction {
    solicitado: string;
    data: string;
    aprovado: string;
    pedido: string;
    saida: string;
    previsao: string;
    entrega: string;
    cidade: string;
    uf: string;
    pontuado: string;
    status: string;
    pdv?: string; // Nome da Loja
    representative_uuid?: string; // Representante responsável
    // Campos adicionais para formulário detalhado
    contact?: string;
    phone?: string;
    email?: string;
    proposal?: string;
    purchaseVolume?: string;
    category?: string; // Categoria da ação (ex: Brindes, Propaganda)
}

export interface GiftItem {
    name: string;
    stock: number;
    price: number;
}

export interface GiftTransaction {
    id: string;
    itemName: string;
    date: string;
    time: string;
    quantity: number;
    unit: string;
    type: 'in' | 'out'; // in = Entrada, out = Saída
    price?: number; // Apenas para entradas
    representative?: string; // Apenas para saídas
}

export interface AccountPayable {
    id: string;
    supplier: string;
    dueDate: string; // YYYY-MM-DD
    description: string;
    amount: number;
    nfArrived: boolean;
    boletoArrived: boolean;
    status: 'pending' | 'sent_to_finance';
    recurrence?: 'monthly' | 'yearly' | 'none';
}

export interface RepresentativeProfile {
    code: number;
    company: string;
    name: string; // Nome Completo / Display Name
    region: string;
    phone: string;
    email: string;
    attendant: string;
    stats?: {
        trainingCount: number;
        onlineCount: number;
        offlineCount: number;
        offlineValue: number;
    };
}

export type PostCategory = 'official' | 'solis_voce' | 'leonardo' | 'luiz';
export type PostType = 'video' | 'static' | 'carousel' | 'story' | 'article_linkedin' | 'article_blog';
export type PostStatus = 'in_progress' | 'review' | 'adjust' | 'approved' | 'published';

export interface PostHistoryEvent {
    id: string;
    action: 'upload' | 'adjust_request' | 'approved' | 'published' | 'status_change';
    user: string;
    text?: string;
    timestamp: string;
}

export interface CalendarPost {
    id: string;
    date: string; // YYYY-MM-DD
    time: string; // HH:MM
    title: string;
    description?: string;
    caption: string;
    category: PostCategory;
    type: PostType; // Novo campo Tipo
    status: PostStatus;
    image_url?: string; // Renomeado de image para image_url (compatível com backend)
    platforms?: string[];
    published_platforms?: string[]; // Renomeado de publishedPlatforms para published_platforms (compatível com backend)
    assignee_id?: string; // ID do responsável (compatível com backend)
    assignee?: PublicUser; // Objeto com dados do responsável (compatível com backend)
    history?: PostHistoryEvent[];
    created_at: string; // Timestamp de criação (compatível com backend)
    updated_at: string; // Timestamp de atualização (compatível com backend)
}

// Tipo específico para Assignee (responsável)
export interface PublicUser {
    id: string;
    name: string;
    email: string;
}

// ============================================
// Social Benchmarking Types (Backend Integration)
// ============================================

// Interface principal (ATUALIZADA - compatível com backend)
export interface SocialBenchmarking {
  id: string;
  brand_name: string;
  avg_likes: number;
  avg_comments: number;
  followers?: number;
  engagement_rate: number;
  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601
}

// Interface para resposta paginada
// NOTA CRÍTICA: O backend usa o campo "benchmarkings" (plural), NÃO "data"
// Verificado em: backend/entrypoint/http/payload/response/social_response.go:16
export interface SocialBenchmarkingListResponse {
  benchmarkings: SocialBenchmarking[];
  meta: {
    total: number;
    page: number;
    limit: number;
    total_pages: number;
  };
}

// Interface para filtros
export interface SocialBenchmarkingFilters {
  brand_name?: string;
  start_date?: string; // YYYY-MM-DD
  end_date?: string;   // YYYY-MM-DD
  active?: boolean;     // true = ativos, false = deletados
  page?: number;
  limit?: number;
  sort_by?: 'engagement_rate' | 'avg_likes' | 'avg_comments' | 'created_at';
  sort_order?: 'asc' | 'desc';
}

// Interface para request de criação
export interface CreateSocialBenchmarkingRequest {
  brand_name: string;
  avg_likes: number;
  avg_comments: number;
  followers?: number;
}

// Interface para request de atualização (partial update)
export interface UpdateSocialBenchmarkingRequest {
  brand_name?: string;
  avg_likes?: number;
  avg_comments?: number;
  followers?: number;
}

export interface ProgramCredential {
    name: string;
    user?: string;
    password?: string;
    access?: string;
    notes?: string;
}

export interface InternalContact {
    name: string;
    role: string;
    department: string;
    extension: string; // Using string as some might have special chars or multiple
    email: string;
}

export interface BudgetItem {
  uuid: string;            // UUID do backend (snake_case)
  cod_obj: string;         // Código do objeto (snake_case)
  obj: string;             // Nome do objeto
  cod_grp: string;         // Código do grupo (snake_case)
  grp: string;             // Nome do grupo
  cod: string;             // Código do item
  desc: string;            // Descrição do item
  vals: number[];          // Array de 12 valores orçados
  realized_vals: number[]; // Array de 12 valores realizados (snake_case)
  year: number;            // Ano do orçamento
  created_at: string;      // Timestamp ISO 8601 (snake_case)
  updated_at: string;      // Timestamp ISO 8601 (snake_case)
  deleted_at?: string;     // Timestamp ISO 8601 (snake_case, opcional)
}

// ============================================
// Calendar Post Types (Backend Integration)
// ============================================

// Interface para filtros de posts do calendário
export interface CalendarFilters {
  category?: string; // Filtrar por categoria (official, solis_voce, leonardo, luiz)
  type?: string; // Filtrar por tipo (video, static, carousel, story, article_linkedin, article_blog)
  status?: string; // Filtrar por status (in_progress, review, adjust, approved, published)
  assignee_id?: string; // Filtrar por responsável
  start_date?: string; // Data inicial (YYYY-MM-DD)
  end_date?: string; // Data final (YYYY-MM-DD)
  platform?: string; // Filtrar por plataforma (instagram, facebook, linkedin, tiktok)
  page?: number; // Número da página
  limit?: number; // Itens por página
  sort_by?: string; // Campo para ordenação
  sort_order?: string; // Ordem (asc, desc)
}

// Interface para resposta paginada do backend
export interface CalendarPostsListResponse {
  data: CalendarPost[]; // Lista de posts
  total: number; // Total de itens
  page: number; // Página atual
  limit: number; // Itens por página
  total_pages: number; // Total de páginas
}

// Interface para request de criação/atualização de post
export interface CalendarPostRequest {
  title: string; // Título do post
  description?: string; // Descrição opcional
  date: string; // Data do post (YYYY-MM-DD)
  time: string; // Hora do post (HH:MM)
  caption?: string; // Caption opcional
  category: string; // Categoria (official, solis_voce, leonardo, luiz)
  type: string; // Tipo (video, static, carousel, story, article_linkedin, article_blog)
  platforms: string[]; // Lista de plataformas
  image_url?: string; // URL da imagem opcional
  assignee_id?: string; // ID do responsável opcional
}

// Interface para request de atualização de status
export interface UpdateStatusRequest {
  status: string; // Novo status (in_progress, review, adjust, approved, published)
  text?: string; // Texto de ajuste (obrigatório quando status = adjust)
}

// ============================================
// Social Post Types (Backend Integration)
// ============================================

// Tipo para plataformas de social media
export type SocialPlatform = 'instagram' | 'facebook' | 'linkedin' | 'tiktok' | 'twitter';

// Tipo para tipos de post de social media
export type SocialPostType = 'image' | 'video' | 'carousel' | 'reel' | 'story';

// Interface principal para SocialPost
export interface SocialPost {
  id: string;
  brand_name: string;
  platform: SocialPlatform;
  post_date: string; // YYYY-MM-DD
  post_time?: string; // HH:MM:SS (opcional)
  likes: number;
  comments: number;
  shares?: number; // Opcional
  post_type: SocialPostType;
  caption?: string; // Opcional
  followers_at_post?: number; // Opcional
  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601
}

// Interface para resposta paginada de SocialPosts
export interface SocialPostListResponse {
  posts: SocialPost[];
  meta: {
    total: number;
    page: number;
    limit: number;
    total_pages: number;
  };
}

// Interface para filtros de SocialPosts
export interface SocialPostFilters {
  brand_name?: string;
  platform?: SocialPlatform;
  post_type?: SocialPostType;
  start_date?: string; // YYYY-MM-DD
  end_date?: string; // YYYY-MM-DD
  min_followers?: number;
  max_followers?: number;
  min_engagement?: number;
  max_engagement?: number;
  active?: boolean; // true = ativos, false = deletados
  page?: number;
  limit?: number;
  sort_by?: 'post_date' | 'likes' | 'comments' | 'shares' | 'created_at';
  sort_order?: 'asc' | 'desc';
}

// Interface para request de criação de SocialPost
export interface CreateSocialPostRequest {
  brand_name: string;
  platform: SocialPlatform;
  post_date: string; // YYYY-MM-DD
  post_time?: string; // HH:MM:SS (opcional)
  likes: number;
  comments: number;
  shares?: number; // Opcional
  post_type: SocialPostType;
  caption?: string; // Opcional
  followers_at_post?: number; // Opcional
}

// Interface para request de atualização de SocialPost (partial update)
export interface UpdateSocialPostRequest {
  brand_name?: string;
  platform?: SocialPlatform;
  post_date?: string; // YYYY-MM-DD
  post_time?: string; // HH:MM:SS (opcional)
  likes?: number;
  comments?: number;
  shares?: number; // Opcional
  post_type?: SocialPostType;
  caption?: string; // Opcional
  followers_at_post?: number; // Opcional
}

// ============================================
// Social Daily Aggregation Types (Backend Integration)
// ============================================

// Interface principal para SocialDailyAggregation
export interface SocialDailyAggregation {
  id: string;
  brand_name: string;
  aggregation_date: string; // YYYY-MM-DD
  total_posts: number;
  total_likes: number;
  total_comments: number;
  total_shares?: number; // Opcional
  avg_likes: number;
  avg_comments: number;
  avg_shares?: number; // Opcional
  followers_at_date?: number; // Opcional
  engagement_rate: number;
  created_at: string; // ISO 8601
  updated_at: string; // ISO 8601
}

// Interface para resposta paginada de SocialDailyAggregations
export interface SocialDailyAggregationListResponse {
  aggregations: SocialDailyAggregation[];
  meta: {
    total: number;
    page: number;
    limit: number;
    total_pages: number;
  };
}

// Interface para filtros de SocialDailyAggregations
export interface SocialDailyAggregationFilters {
  brand_name?: string;
  platform?: SocialPlatform;
  start_date?: string; // YYYY-MM-DD
  end_date?: string; // YYYY-MM-DD
  active?: boolean; // true = ativos, false = deletados
  page?: number;
  limit?: number;
  sort_by?: 'aggregation_date' | 'total_posts' | 'total_likes' | 'total_comments' | 'engagement_rate' | 'created_at';
  sort_order?: 'asc' | 'desc';
}
