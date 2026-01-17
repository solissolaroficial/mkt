
import { KpiCategory, RepTableData, MarketingChannelData, Task, AllowedUser, PdvPost, RecurrentPdv, KpiLog, ShowroomItem, CalendarPost, SocialBenchmarking, Notification, GiftItem, GiftTransaction, AccountPayable, RepresentativeProfile, ProgramCredential, InternalContact, BreakdownItem, OfflineAction, BudgetItem, ChannelPerformance } from '../types/legacy.types';

export const MONTHS = ['JAN', 'FEV', 'MAR', 'ABR', 'MAI', 'JUN', 'JUL', 'AGO', 'SET', 'OUT', 'NOV', 'DEZ'];

export const FULL_MONTH_MAP: Record<string, string> = {
  'JAN': 'JANEIRO', 'FEV': 'FEVEREIRO', 'MAR': 'MARÇO', 'ABR': 'ABRIL',
  'MAI': 'MAIO', 'JUN': 'JUNHO', 'JUL': 'JULHO', 'AGO': 'AGOSTO',
  'SET': 'SETEMBRO', 'OUT': 'OUTUBRO', 'NOV': 'NOVEMBRO', 'DEZ': 'DEZEMBRO',
  '---': 'ANO COMPLETO'
};

export const ALLOWED_USERS: AllowedUser[] = ['Jackson', 'Beatriz', 'Larissa'];

// Helper to parse "1.000" or "-" to number
const parseVal = (v: string): number => {
  if (!v || v === '-') return 0;
  return parseFloat(v.replace(/\./g, '').replace(',', '.'));
};

// Helper to create item with simulated realized data
const item = (c1: string, c2: string, c3: string, c4: string, c5: string, c6: string, v: string[]): BudgetItem => {
  const parsedVals = v.map(parseVal);
  // Simulating realized values for demo
  const realizedVals = parsedVals.map((val, idx) => {
      if (val === 0) return 0;
      // Simulate partial realization for first few months
      if (idx < 3) {
          const variance = (Math.random() * 0.4) - 0.1; // -10% to +30%
          return Math.floor(val * (1 + variance));
      }
      return 0; 
  });

  return {
    id: Math.random().toString(36).substr(2, 9),
    codObj: c1, obj: c2, codGrp: c3, grp: c4, cod: c5, desc: c6,
    vals: parsedVals,
    realizedVals: realizedVals
  };
};

export const RAW_BUDGET_ITEMS: BudgetItem[] = [
  item("9", "EVENTOS & REUNIÕES", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS E PUBLICIDADES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "20", "DESPESAS C/ MANUTENÇÃO", "85", "PEÇAS E ACESSÓRIOS", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "89", "MATERIAIS DE EXPEDIENTE E CONSUMO", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "87", "SERVIÇOS DE TERCEIROS - PESSOA JURÍDICA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "125", "SERVIÇOS DE FRETES E TRANSPORTES", ["-", "-", "-", "7.000", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "31", "DESPESAS C/ CURSOS & TREINAMENTOS", "160", "DESPESAS C/ CURSO/TREINAMENTO (INTERNO)", ["-", "-", "-", "2.500", "-", "2.500", "-", "2.500", "2.500", "2.500", "-", "12.500"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "528", "PALESTRAS/CURSOS/TREINAMENTOS - AÇÃO COMERCIAL", ["-", "5.000", "5.000", "-", "5.000", "-", "3.000", "3.000", "5.000", "3.000", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON", ["8.235", "8.235", "8.235", "8.235", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON ESTANDE", ["30.000", "30.000", "30.000", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON HOSPEDAGEM", ["-", "-", "-", "8.000", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON ALIMENTAÇÃO", ["-", "-", "-", "5.500", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON COMBUSTÍVEL", ["-", "-", "-", "3.500", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON TAXAS", ["-", "-", "-", "1.000", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | FEICON BEBIDAS", ["-", "-", "-", "750", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | BIRIGUI NELSON WORKSHOP", ["-", "-", "5.000", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | ANIVERSÁRIO DA SOLIS", ["-", "-", "-", "-", "40.000", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | CONVENÇÃO DE VENDAS", ["-", "-", "-", "-", "30.000", "-", "-", "-", "-", "-", "-", "-"]),
  item("9", "EVENTOS & REUNIÕES", "41", "DESPESAS COMERCIAIS", "537", "FEIRAS | AECIA", ["1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "-", "-", "-", "-", "-", "-"]),
  item("10", "PROMOÇÕES LOCAIS", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "380", "OUTDOOR", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "13", "DESPESAS SOCIAIS", "56", "CONFRATERNIZAÇÕES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS E PUBLICIDADES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "87", "SERVIÇOS DE TERCEIROS - PESSOA JURÍDICA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "125", "SERVIÇOS DE FRETES E TRANSPORTES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "41", "DESPESAS COMERCIAIS", "395", "AMOSTRAS - PDV", ["16.000", "16.000", "16.000", "16.000", "16.000", "16.000", "16.000", "16.000", "16.000", "-", "-", "-"]),
  item("12", "BRINDES", "41", "DESPESAS COMERCIAIS", "395", "BRINDES - AÇÃO COMERCIAL", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("12", "BRINDES", "41", "DESPESAS COMERCIAIS", "525", "BRINDES - COOPERADA", ["-", "12.000", "12.000", "12.000", "-", "12.000", "12.000", "12.000", "12.000", "12.000", "-", "-"]),
  item("12", "BRINDES", "47", "OUTRAS DESPESAS COM FUNCIONÁRIO", "92", "UNIFORMES", ["-", "2.500", "2.500", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "14", "DESPESAS C/ VIAGENS", "139", "DESPESAS DE VIAGEM", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | ZECCA CONSONI", ["8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000"]),
  item("14", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | PLOTAGEM CAMINHÃO", ["10.000", "10.000", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("15", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | GRATITUDE", ["5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500", "5.500"]),
  item("16", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | ANIMAÇÃO 3D", ["1.200", "2.030", "1.200", "1.700", "-", "-", "1.600", "1.600", "-", "-", "-", "-"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | RUBENS", ["450", "450", "450", "450", "450", "450", "450", "450", "450", "450", "450", "450"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | RD STATION", ["1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345", "1.345"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | CURSEDUCA", ["1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000", "1.000"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | PORTAL DE NOTÍCIAS", ["1.000", "1.000", "8.000", "1.000", "1.000", "1.000", "-", "-", "-", "-", "-", "-"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | BIM", ["2.000", "2.000", "1.000", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("17", "PROPAGANDAS & PUBLICIDADE", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS & PUBLICIDADE | VÍDEO INSTITUCIONAL", ["10.000", "10.000", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "88", "SERVIÇO REFEIÇÕES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "89", "MATERIAIS DE EXPEDIENTE E CONSUMO", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "87", "SERVIÇOS DE TERCEIROS - PESSOA JURÍDICA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "41", "DESPESAS COMERCIAIS", "164", "DESPESAS C/ PROPAGANDA/PUBLIC", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "41", "DESPESAS COMERCIAIS", "524", "BRINDES - AÇÃO COMERCIAL", ["20.000", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "10.000", "-", "46.000"]),
  item("13", "PROPAGANDAS & PUBLICIDADE", "41", "DESPESAS COMERCIAIS", "532", "MKT DIGITAL", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("14", "MÍDIA DIGITAL", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS E PUBLICIDADES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("14", "MÍDIA DIGITAL", "41", "DESPESAS COMERCIAIS", "164", "DESPESAS C/ PROPAGANDA/PUBLIC (INFLUENCIADOR)", ["8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000", "8.000"]),
  item("14", "MÍDIA DIGITAL", "41", "DESPESAS COMERCIAIS", "369", "SITE", ["5.033", "5.033", "5.033", "-", "-", "500", "-", "-", "-", "-", "-", "-"]),
  item("14", "MÍDIA DIGITAL", "41", "DESPESAS COMERCIAIS", "371", "GOOGLE", ["-", "3.500", "4.000", "4.000", "4.000", "4.000", "4.000", "4.000", "4.000", "4.000", "4.000", "-"]),
  item("15", "MÍDIA TRADICIONAL", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "87", "SERVIÇOS DE TERCEIROS - PESSOA JURÍDICA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("15", "MÍDIA TRADICIONAL", "41", "DESPESAS COMERCIAIS", "164", "DESPESAS C/ PROPAGANDA/PUBLIC", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("15", "MÍDIA TRADICIONAL", "41", "DESPESAS COMERCIAIS", "524", "BRINDES - AÇÃO COMERCIAL", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("15", "MÍDIA TRADICIONAL", "41", "DESPESAS COMERCIAIS", "526", "IMPRESSOS", ["5.000", "5.000", "5.000", "5.000", "150", "150", "150", "150", "150", "150", "150", "10.000"]),
  item("341", "APOIO MARKETING", "13", "DESPESAS SOCIAIS", "56", "CONFRATERNIZAÇÕES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "14", "DESPESAS C/ VIAGENS", "139", "DESPESAS DE VIAGEM", ["-", "-", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "2.000", "-", "-"]),
  item("341", "APOIO MARKETING", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "64", "TELEFONE E INTERNET", ["574", "294", "280", "-", "559", "-", "245", "279", "559", "277", "240", "240"]),
  item("341", "APOIO MARKETING", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "65", "CORREIOS", ["150", "150", "150", "150", "150", "150", "150", "150", "150", "150", "50", "50"]),
  item("341", "APOIO MARKETING", "15", "DESPESAS C/ COMUNICAÇÃO E INFORMAÇÕES", "67", "PROPAGANDAS E PUBLICIDADES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "16", "DESPESAS LEGAIS - IMPOSTOS/TAXAS", "148", "ASSOCIAÇÃO DE CLASSE", ["-", "-12", "-12", "-18", "-18", "-18", "-18", "-18", "-18", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "20", "DESPESAS C/ MANUTENÇÃO", "85", "PEÇAS E ACESSÓRIOS", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "20", "DESPESAS C/ MANUTENÇÃO", "142", "MATERIAIS DE CONSTRUÇÃO", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "88", "SERVIÇO REFEIÇÕES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "89", "MATERIAIS DE EXPEDIENTE E CONSUMO", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | POWERBANK", ["-", "-", "-", "-", "-", "250", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | NOTEBOOK", ["-", "-", "-", "-", "-", "7.000", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | CABO TETHERING", ["-", "300", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | CARTÃO DE MEMÓRIA", ["-", "150", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | MONITOR", ["-", "-", "-", "-", "-", "1.900", "-", "1.900", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "90", "PERIFÉRICOS DE INFORMÁTICA | MICROFONE E TRIPÉ", ["3.000", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "21", "DESPESAS C/ MATERIAIS E CONSUMO", "91", "MATERIAIS DE SEGURANÇA - EPI", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "24", "MATÉRIA PRIMA", "100", "MATÉRIA PRIMA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "87", "SERVIÇOS DE TERCEIROS - PESSOA JURÍDICA", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "125", "SERVIÇOS DE FRETES E TRANSPORTES", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "344", "LICENÇA DE USO | CHAT GPT", ["100", "100", "100", "100", "100", "100", "100", "100", "100", "100", "100", "100"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "344", "LICENÇA DE USO | ADOBE", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "3.000", "-", "-"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "344", "LICENÇA DE USO | OPUS", ["100", "100", "100", "100", "100", "100", "100", "100", "100", "100", "100", "100"]),
  item("341", "APOIO MARKETING", "25", "DESPESAS C/ SERVIÇOS DE TERCEIROS", "344", "LICENÇA DE USO | CAPCUT", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "370", "-", "-"]),
  item("341", "APOIO MARKETING", "31", "DESPESAS C/ CURSOS & TREINAMENTOS", "160", "DESPESAS C/ CURSO/TREINAMENTO", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "41", "DESPESAS COMERCIAIS", "367", "TREINAMENTO PRESENCIAL", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "41", "DESPESAS COMERCIAIS", "368", "WORKSHOPS", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "41", "DESPESAS COMERCIAIS", "524", "BRINDES - AÇÃO COMERCIAL", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "44", "RECEITAS DIVERSAS", "126", "RECUPERAÇÃO DE CUSTOS\\DESPESAS", ["-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-", "-"]),
  item("341", "APOIO MARKETING", "47", "OUTRAS DESPESAS COM FUNCIONÁRIO", "222", "DESPESAS C/ VALE REFEIÇÃO/CARTÃO", ["420", "420", "420", "420", "420", "420", "420", "420", "420", "420", "900", "900"]),
];

const generateInitialLogs = (data: { month: string; realized: number | null }[]): KpiLog[] => {
  return data
    .filter(d => d.realized !== null)
    .map((d, index) => {
      const monthIndex = MONTHS.indexOf(d.month);
      const date = new Date(2025, monthIndex, 1, 9, 0, 0).toISOString();
      
      return {
        id: `init-${d.month}-${index}`,
        date: date,
        timestamp: '09:00',
        user: 'Sistema',
        month: d.month,
        oldValue: null,
        newValue: d.realized!,
        action: 'create' as const,
        context: 'Carga Inicial'
      };
    })
    .reverse(); 
};

const generateGenericBreakdown = (total: number): BreakdownItem[] => {
    const site = Math.floor(total * 0.45);
    const ads = total - site;
    const br = Math.floor(ads * 0.6);
    const rj = Math.floor(ads * 0.15);
    const mg = Math.floor(ads * 0.10);
    const es = ads - br - rj - mg;

    return [
        {
            label: 'Campanhas Ads Meta (Instagram/Facebook)',
            value: ads,
            subItems: [
                { label: 'Brasil', value: br },
                { label: 'RJ', value: rj },
                { label: 'Sul de MG', value: mg },
                { label: 'Espírito Santo', value: es }
            ]
        },
        {
            label: 'Botão no site',
            value: site
        }
    ];
};

const generateGenericAdBreakdown = (total: number): BreakdownItem[] => {
    const adsMeta = Math.floor(total * 0.60); 
    const googleAds = total - adsMeta; 
    
    const br = Math.floor(adsMeta * 0.5);
    const rj = Math.floor(adsMeta * 0.2);
    const mg = Math.floor(adsMeta * 0.15);
    const es = adsMeta - br - rj - mg;

    return [
        {
            label: 'Campanhas Ads Meta (Instagram/Facebook)',
            value: adsMeta,
            subItems: [
                { label: 'Brasil', value: br },
                { label: 'RJ', value: rj },
                { label: 'Sul de MG', value: mg },
                { label: 'Espírito Santo', value: es }
            ]
        },
        {
            label: 'Botão no site',
            value: googleAds
        }
    ];
};

const generateGenericRoasBreakdown = (totalRoas: number): BreakdownItem[] => {
    return [
        {
            label: 'Campanhas Ads Meta (Instagram/Facebook)',
            value: Number((totalRoas * 1.1).toFixed(2)) 
        },
        {
            label: 'Botão no site', 
            value: Number((totalRoas * 0.85).toFixed(2)) 
        },
        {
            label: 'E-mail Marketing',
            value: Number((totalRoas * 1.5).toFixed(2)) 
        }
    ];
};

export const MOCK_NOTIFICATIONS: Notification[] = [
    {
        id: 'n1',
        type: 'deadline',
        title: 'Tarefa Vencendo',
        message: 'A tarefa "Reunião com representantes" vence hoje.',
        timestamp: 'Há 2 horas',
        read: false,
        taskId: '2'
    },
    {
        id: 'n2',
        type: 'mention',
        title: 'Nova Menção',
        message: 'Beatriz mencionou você em "Analisar relatório de leads".',
        timestamp: 'Há 4 horas',
        read: false,
        taskId: '1'
    },
    {
        id: 'n3',
        type: 'system',
        title: 'Atualização do Sistema',
        message: 'Novos dados de benchmarking social disponíveis.',
        timestamp: 'Ontem',
        read: true
    }
];

const RAW_KPIS: KpiCategory[] = [
  {
    id: 'youtube',
    title: 'Novos Inscritos no Youtube',
    slug: 'novos_inscritos_no_youtube',
    color: '#FF0000',
    logs: [],
    data: [
      { month: 'JAN', realized: 14, meta: 15 },
      { month: 'FEV', realized: 66, meta: 25 },
      { month: 'MAR', realized: 270, meta: 25 },
      { month: 'ABR', realized: 81, meta: 50 },
      { month: 'MAI', realized: 57, meta: 75 },
      { month: 'JUN', realized: 45, meta: 75 },
      { month: 'JUL', realized: 55, meta: 100 },
      { month: 'AGO', realized: 66, meta: 120 },
      { month: 'SET', realized: 33, meta: 120 },
      { month: 'OUT', realized: 66, meta: 115 },
      { month: 'NOV', realized: 45, meta: 75 },
      { month: 'DEZ', realized: null, meta: 45 },
    ]
  },
  {
    id: 'training_reps',
    title: 'Treinamentos Realizados (Reps)',
    slug: 'treinamentos_realizados_reps',
    color: '#3B82F6',
    logs: [],
    data: [
      { month: 'JAN', realized: 0, meta: 0 },
      { month: 'FEV', realized: 3, meta: 15 },
      { month: 'MAR', realized: 5, meta: 30 },
      { month: 'ABR', realized: 10, meta: 30 },
      { month: 'MAI', realized: 5, meta: 30 },
      { month: 'JUN', realized: 12, meta: 30 },
      { month: 'JUL', realized: 23, meta: 30 },
      { month: 'AGO', realized: 22, meta: 30 },
      { month: 'SET', realized: 22, meta: 30 },
      { month: 'OUT', realized: 21, meta: 30 },
      { month: 'NOV', realized: 11, meta: 15 },
      { month: 'DEZ', realized: null, meta: 15 },
    ]
  },
  {
    id: 'training_people',
    title: 'Pessoas Treinadas (Solis)',
    slug: 'pessoas_treinadas_solis',
    color: '#10B981',
    logs: [],
    data: [
      { month: 'JAN', realized: 5, meta: 0 },
      { month: 'FEV', realized: 44, meta: 30 },
      { month: 'MAR', realized: 38, meta: 40 },
      { month: 'ABR', realized: 0, meta: 20 },
      { month: 'MAI', realized: 77, meta: 55 },
      { month: 'JUN', realized: 43, meta: 65 },
      { month: 'JUL', realized: 97, meta: 65 },
      { month: 'AGO', realized: 85, meta: 55 },
      { month: 'SET', realized: 24, meta: 20 },
      { month: 'OUT', realized: 58, meta: 50 },
      { month: 'NOV', realized: 56, meta: 40 },
      { month: 'DEZ', realized: null, meta: 30 },
    ]
  },
  {
    id: 'marketing',
    title: 'Ações de Marketing (Reps)',
    slug: 'acoes_de_marketing_reps',
    color: '#8B5CF6',
    logs: [],
    data: [
      { month: 'JAN', realized: 1, meta: 0 },
      { month: 'FEV', realized: 10, meta: 30 },
      { month: 'MAR', realized: 13, meta: 60 },
      { month: 'ABR', realized: 82, meta: 60 },
      { month: 'MAI', realized: 18, meta: 60 },
      { month: 'JUN', realized: 33, meta: 60 },
      { month: 'JUL', realized: 48, meta: 60 },
      { month: 'AGO', realized: 64, meta: 60 },
      { month: 'SET', realized: 62, meta: 60 },
      { month: 'OUT', realized: 48, meta: 60 },
      { month: 'NOV', realized: 35, meta: 30 },
      { month: 'DEZ', realized: null, meta: 30 },
    ]
  },
  {
    id: 'inbound',
    title: 'Visitantes Inbound',
    slug: 'visitantes_inbound',
    color: '#F59E0B',
    logs: [],
    data: [
      { month: 'JAN', realized: 3417, meta: 12500 },
      { month: 'FEV', realized: 3266, meta: 14000 },
      { month: 'MAR', realized: 3437, meta: 19000 },
      { month: 'ABR', realized: 3396, meta: 19000 },
      { month: 'MAI', realized: 3604, meta: 18000 },
      { month: 'JUN', realized: 3838, meta: 18000 },
      { month: 'JUL', realized: 4590, meta: 18000 },
      { month: 'AGO', realized: 4650, meta: 18000 },
      { month: 'SET', realized: 3968, meta: 18000 },
      { month: 'OUT', realized: 4473, meta: 18000 },
      { month: 'NOV', realized: 3201, meta: 15000 },
      { month: 'DEZ', realized: null, meta: 12500 },
    ]
  },
  {
    id: 'opportunities',
    title: 'Taxa de Oportunidades',
    slug: 'taxa_de_oportunidades',
    color: '#EC4899',
    unit: 'number',
    logs: [],
    data: [
      { month: 'JAN', realized: 147, meta: 100, breakdown: generateGenericBreakdown(147) },
      { month: 'FEV', realized: 69, meta: 100, breakdown: generateGenericBreakdown(69) },
      { month: 'MAR', realized: 270, meta: 170, breakdown: generateGenericBreakdown(270) },
      { month: 'ABR', realized: 778, meta: 200, breakdown: generateGenericBreakdown(778) },
      { month: 'MAI', realized: 160, meta: 170, breakdown: generateGenericBreakdown(160) },
      { month: 'JUN', realized: 152, meta: 250, breakdown: generateGenericBreakdown(152) },
      { month: 'JUL', realized: 224, meta: 250, breakdown: generateGenericBreakdown(224) },
      { month: 'AGO', realized: 278, meta: 200, breakdown: generateGenericBreakdown(278) },
      { month: 'SET', realized: 324, meta: 180, breakdown: generateGenericBreakdown(324) },
      { month: 'OUT', realized: 392, meta: 300, breakdown: generateGenericBreakdown(392) },
      { 
        month: 'NOV', 
        realized: 344, 
        meta: 200,
        breakdown: [
            {
                label: 'Campanhas Ads Meta (Instagram/Facebook)',
                value: 178,
                subItems: [
                    { label: 'Brasil', value: 118 },
                    { label: 'RJ', value: 26 },
                    { label: 'Sul de MG', value: 11 },
                    { label: 'Espírito Santo', value: 23 }
                ]
            },
            {
                label: 'Botão no site',
                value: 166
            }
        ]
      },
      { month: 'DEZ', realized: null, meta: 80 },
    ]
  },
  {
    id: 'authority',
    title: 'Autoridade na Internet (DA)',
    slug: 'autoridade_na_internet_da',
    color: '#6366F1',
    logs: [],
    data: [
      { month: 'JAN', realized: 30, meta: 30 },
      { month: 'FEV', realized: 30, meta: 30 },
      { month: 'MAR', realized: 30, meta: 30 },
      { month: 'ABR', realized: 30, meta: 30 },
      { month: 'MAI', realized: 33, meta: 31 },
      { month: 'JUN', realized: 33, meta: 31 },
      { month: 'JUL', realized: 33, meta: 33 },
      { month: 'AGO', realized: 35, meta: 33 },
      { month: 'SET', realized: 37, meta: 33 },
      { month: 'OUT', realized: 38, meta: 33 },
      { month: 'NOV', realized: 38, meta: 33 },
      { month: 'DEZ', realized: 38, meta: 33 },
    ]
  },
  {
    id: 'ad_spend',
    title: 'Investimento em Ads (R$)',
    slug: 'investimento_em_ads',
    color: '#85bb65',
    unit: 'currency',
    logs: [],
    data: [
      { month: 'JAN', realized: 12500, meta: 12000, breakdown: generateGenericAdBreakdown(12500) },
      { month: 'FEV', realized: 13200, meta: 13000, breakdown: generateGenericAdBreakdown(13200) },
      { month: 'MAR', realized: 15800, meta: 15000, breakdown: generateGenericAdBreakdown(15800) },
      { month: 'ABR', realized: 14900, meta: 15000, breakdown: generateGenericAdBreakdown(14900) },
      { month: 'MAI', realized: 18500, meta: 18000, breakdown: generateGenericAdBreakdown(18500) },
      { month: 'JUN', realized: 19200, meta: 20000, breakdown: generateGenericAdBreakdown(19200) },
      { month: 'JUL', realized: 21500, meta: 20000, breakdown: generateGenericAdBreakdown(21500) },
      { month: 'AGO', realized: 22000, meta: 20000, breakdown: generateGenericAdBreakdown(22000) },
      { month: 'SET', realized: 19800, meta: 20000, breakdown: generateGenericAdBreakdown(19800) },
      { month: 'OUT', realized: 23500, meta: 22000, breakdown: generateGenericAdBreakdown(23500) },
      { 
        month: 'NOV', 
        realized: 18000, 
        meta: 20000,
        breakdown: [
            {
                label: 'Campanhas Ads Meta (Instagram/Facebook)',
                value: 12500,
                subItems: [
                    { label: 'Brasil', value: 7000 },
                    { label: 'RJ', value: 2500 },
                    { label: 'Sul de MG', value: 1000 },
                    { label: 'Espírito Santo', value: 2000 }
                ]
            },
            {
                label: 'Botão no site',
                value: 5500
            }
        ]
      },
      { month: 'DEZ', realized: null, meta: 15000 },
    ]
  },
  {
    id: 'cpl',
    title: 'Custo por Lead (CPL)',
    slug: 'custo_por_lead_cpl',
    color: '#eab308',
    unit: 'currency',
    logs: [],
    data: [
      { month: 'JAN', realized: 45, meta: 40 },
      { month: 'FEV', realized: 42, meta: 40 },
      { month: 'MAR', realized: 38, meta: 38 },
      { month: 'ABR', realized: 35, meta: 38 },
      { month: 'MAI', realized: 39, meta: 35 },
      { month: 'JUN', realized: 41, meta: 35 },
      { month: 'JUL', realized: 36, meta: 35 },
      { month: 'AGO', realized: 34, meta: 35 },
      { month: 'SET', realized: 37, meta: 35 },
      { month: 'OUT', realized: 33, meta: 32 },
      { month: 'NOV', realized: 35, meta: 32 },
      { month: 'DEZ', realized: null, meta: 32 },
    ]
  },
  {
    id: 'mql',
    title: 'Leads Qualificados (MQLs)',
    slug: 'leads_qualificados_mqls',
    color: '#06b6d4',
    logs: [],
    data: [
      { month: 'JAN', realized: 85, meta: 100 },
      { month: 'FEV', realized: 110, meta: 120 },
      { month: 'MAR', realized: 145, meta: 150 },
      { month: 'ABR', realized: 180, meta: 160 },
      { month: 'MAI', realized: 210, meta: 200 },
      { month: 'JUN', realized: 195, meta: 200 },
      { month: 'JUL', realized: 230, meta: 220 },
      { month: 'AGO', realized: 255, meta: 240 },
      { month: 'SET', realized: 240, meta: 240 },
      { month: 'OUT', realized: 280, meta: 260 },
      { month: 'NOV', realized: 190, meta: 200 },
      { month: 'DEZ', realized: null, meta: 150 },
    ]
  },
  {
    id: 'roas',
    title: 'ROAS (Retorno em Ads)',
    slug: 'roas_retorno_em_ads',
    color: '#6366f1',
    unit: 'number',
    logs: [],
    data: [
      { month: 'JAN', realized: 3.2, meta: 4.0, breakdown: generateGenericRoasBreakdown(3.2) },
      { month: 'FEV', realized: 3.5, meta: 4.0, breakdown: generateGenericRoasBreakdown(3.5) },
      { month: 'MAR', realized: 4.2, meta: 4.5, breakdown: generateGenericRoasBreakdown(4.2) },
      { month: 'ABR', realized: 4.8, meta: 4.5, breakdown: generateGenericRoasBreakdown(4.8) },
      { month: 'MAI', realized: 4.5, meta: 5.0, breakdown: generateGenericRoasBreakdown(4.5) },
      { month: 'JUN', realized: 4.1, meta: 5.0, breakdown: generateGenericRoasBreakdown(4.1) },
      { month: 'JUL', realized: 5.2, meta: 5.0, breakdown: generateGenericRoasBreakdown(5.2) },
      { month: 'AGO', realized: 5.5, meta: 5.5, breakdown: generateGenericRoasBreakdown(5.5) },
      { month: 'SET', realized: 5.1, meta: 5.5, breakdown: generateGenericRoasBreakdown(5.1) },
      { month: 'OUT', realized: 5.8, meta: 6.0, breakdown: generateGenericRoasBreakdown(5.8) },
      { 
        month: 'NOV', 
        realized: 4.9, 
        meta: 6.0,
        breakdown: [
            { label: 'Campanhas Ads Meta (Instagram/Facebook)', value: 5.2 },
            { label: 'Botão no site', value: 4.1 },
            { label: 'E-mail Marketing', value: 8.5 }
        ]
      },
      { month: 'DEZ', realized: null, meta: 5.0 },
    ]
  },
  {
    id: 'conversion_rate',
    title: 'Taxa de Conversão Global (%)',
    slug: 'taxa_de_conversao_global',
    color: '#f43f5e',
    unit: 'percent',
    logs: [],
    data: [
      { month: 'JAN', realized: 2.1, meta: 2.5 },
      { month: 'FEV', realized: 2.3, meta: 2.5 },
      { month: 'MAR', realized: 2.8, meta: 3.0 },
      { month: 'ABR', realized: 3.2, meta: 3.0 },
      { month: 'MAI', realized: 3.0, meta: 3.5 },
      { month: 'JUN', realized: 2.9, meta: 3.5 },
      { month: 'JUL', realized: 3.5, meta: 3.5 },
      { month: 'AGO', realized: 3.8, meta: 4.0 },
      { month: 'SET', realized: 3.4, meta: 4.0 },
      { month: 'OUT', realized: 4.1, meta: 4.0 },
      { month: 'NOV', realized: 3.2, meta: 4.0 },
      { month: 'DEZ', realized: null, meta: 3.5 },
    ]
  }
];

export const KPIS: KpiCategory[] = RAW_KPIS.map(kpi => ({
  ...kpi,
  logs: generateInitialLogs(kpi.data)
}));

export const REP_NAMES = [
  'André', 'César', 'Cleber', 'Cristiano e Ranoika', 'Fausto', 'Gonçalves', 
  'Márcio Henrique', 'Marcos', 'Nilton', 'Jorge', 'Otávio', 'Rafael Betoni', 
  'Rafael Lazzarotto', 'Wilson'
];

export const REP_PROFILES: Record<string, RepresentativeProfile> = {
    'André': {
        code: 466,
        company: 'MELO & DOMINGUES REP',
        name: 'André',
        region: 'Nordeste',
        phone: '(81) 98268-3570\n(81) 3221-6632\n(81) 3203-4036',
        email: 'andremelo@solis.ind.br',
        attendant: 'Fran'
    },
};

export const MOCK_TASKS: Task[] = [
  {
    id: '1',
    title: 'Analisar relatório de leads de Outubro',
    description: 'Verificar a qualidade dos leads vindos do Google Ads e comparar com o mês anterior. Identificar quais campanhas trouxeram mais MQLs.',
    due_date: '2025-11-20T12:00:00Z',
    priority: 'high',
    status: 'pending',
    category: 'marketing',
    assignee_id: 'Jackson',
    flows: ['Jackson'], // Added flows
    archived: false,
    subtasks: [
      { id: '1-1', title: 'Exportar dados do CRM', completed: true, assignee: 'Jackson', dueDate: '2025-11-21' },
      { id: '1-2', title: 'Cruzar com dados do Google Ads', completed: false, assignee: 'Jackson', dueDate: '2025-11-21' },
      { id: '1-3', title: 'Montar apresentação para diretoria', completed: false, assignee: 'Beatriz', dueDate: '2025-11-22' }
    ],
    comments: [
      { id: 'c1', user: 'Beatriz', text: 'Já exportei a base bruta, está na pasta compartilhada.', timestamp: '10:30' }
    ]
  },
  {
    id: '2',
    title: 'Reunião com representantes regionais',
    description: 'Alinhamento mensal de metas e novidades do portfólio.',
    due_date: '2025-11-21T12:00:00Z',
    priority: 'medium',
    status: 'in_progress',
    category: 'commercial',
    assignee_id: 'Larissa',
    flows: ['Larissa'], // Added flows
    archived: false,
    subtasks: [
      { id: '2-1', title: 'Preparar pauta', completed: true, assignee: 'Larissa', dueDate: '2025-11-20' },
      { id: '2-2', title: 'Enviar convites', completed: true, assignee: 'Larissa', dueDate: '2025-11-20' }
    ],
    comments: []
  },
  {
    id: '3',
    title: 'Aprovar criativos da campanha de Black Friday',
    description: 'Revisar textos e imagens para a campanha de fim de ano.',
    due_date: '2025-11-21T12:00:00Z',
    priority: 'high',
    status: 'completed',
    category: 'marketing',
    assignee_id: 'Beatriz',
    flows: ['Beatriz', 'Jackson'], // Multi-flow example
    archived: false,
    subtasks: [
        { id: '3-1', title: 'Revisão ortográfica', completed: true, assignee: 'Larissa', dueDate: '2025-11-20' },
        { id: '3-2', title: 'Aprovação da diretoria', completed: true, assignee: 'Beatriz', dueDate: '2025-11-21' }
    ],
    comments: []
  },
  {
    id: '4',
    title: 'Atualizar planilha de metas Q4',
    description: 'Inserir os dados consolidados do Q3 e projetar Q4.',
    due_date: '2025-11-22T12:00:00Z',
    priority: 'high',
    status: 'pending',
    category: 'administrative',
    assignee_id: 'Jackson',
    flows: ['Jackson'],
    archived: false,
    subtasks: [],
    comments: []
  },
  {
    id: '5',
    title: 'Feedback trimestral equipe de vendas',
    description: 'Rodada de feedbacks individuais.',
    due_date: '2025-11-28T12:00:00Z',
    priority: 'medium',
    status: 'pending',
    category: 'commercial',
    assignee_id: 'Jackson',
    flows: ['Jackson'],
    archived: false,
    subtasks: [],
    comments: []
  },
  {
    id: '6',
    title: 'Renovar domínio do site',
    description: 'Pagamento anual do registro.br',
    due_date: '2025-11-28T12:00:00Z',
    priority: 'low',
    status: 'pending',
    category: 'administrative',
    assignee_id: 'Larissa',
    flows: ['Larissa'],
    archived: false,
    subtasks: [],
    comments: []
  },
  {
    id: '7',
    title: 'Planejamento estratégico 2026',
    description: 'Início das reuniões de planejamento.',
    due_date: '2025-11-28T12:00:00Z',
    priority: 'high',
    status: 'in_progress',
    category: 'administrative',
    assignee_id: 'Beatriz',
    flows: ['Beatriz'],
    archived: false,
    subtasks: [
        { id: '7-1', title: 'Definir datas', completed: true, assignee: 'Beatriz', dueDate: '2025-11-21' },
        { id: '7-2', title: 'Reservar sala', completed: false, assignee: 'Larissa', dueDate: '2025-11-22' }
    ],
    comments: []
  },
  {
    id: '8',
    title: 'teste',
    description: 'teste',
    due_date: '2025-12-16T12:00:00Z',
    priority: 'medium',
    status: 'pending',
    category: 'marketing',
    assignee_id: 'Jackson',
    flows: ['Jackson'],
    subtasks: [],
    comments: [],
    archived: false
  }
];

export const REPS_TRAINING_DATA: RepTableData = {
  title: 'Histórico de Treinamentos',
  columns: REP_NAMES,
  rows: MONTHS.map(month => ({
    month: FULL_MONTH_MAP[month],
    target: 2, // Example target
    values: REP_NAMES.reduce((acc, rep) => ({ ...acc, [rep]: Math.floor(Math.random() * 4) }), {})
  }))
};

export const REPS_MARKETING_DATA: RepTableData = {
  title: 'Ações de Marketing por Representante',
  columns: REP_NAMES,
  rows: MONTHS.map(month => ({
    month: FULL_MONTH_MAP[month],
    target: 1, // Example target
    values: REP_NAMES.reduce((acc, rep) => ({ ...acc, [rep]: Math.floor(Math.random() * 2) }), {})
  }))
};

export const SOCIAL_MEDIA_RANKING: SocialBenchmarking[] = [
    {
        id: '1',
        brand_name: 'Solis Solar',
        avg_likes: 145.5,
        avg_comments: 12.3,
        followers: 15400,
        engagement_rate: 1.03,
        created_at: '2025-01-01T00:00:00Z',
        updated_at: '2025-01-01T00:00:00Z'
    },
    {
        id: '2',
        brand_name: 'Competitor A',
        avg_likes: 120.2,
        avg_comments: 8.5,
        followers: 12000,
        engagement_rate: 1.07,
        created_at: '2025-01-01T00:00:00Z',
        updated_at: '2025-01-01T00:00:00Z'
    },
    {
        id: '3',
        brand_name: 'Competitor B',
        avg_likes: 98.4,
        avg_comments: 5.2,
        followers: 8500,
        engagement_rate: 1.22,
        created_at: '2025-01-01T00:00:00Z',
        updated_at: '2025-01-01T00:00:00Z'
    },
    {
        id: '4',
        brand_name: 'Competitor C',
        avg_likes: 210.8,
        avg_comments: 18.9,
        followers: 25000,
        engagement_rate: 0.92,
        created_at: '2025-01-01T00:00:00Z',
        updated_at: '2025-01-01T00:00:00Z'
    },
];

export const MOCK_CALENDAR_POSTS: CalendarPost[] = [
    {
        id: '1',
        date: '2025-11-20',
        time: '10:00',
        title: 'Lançamento Linha 2026',
        description: 'Post sobre os novos coletores de alta eficiência.',
        caption: 'Chegou a hora de conhecer o futuro da energia solar! ☀️🚀 #SolisSolar #EnergiaSolar',
        category: 'official',
        type: 'video',
        status: 'approved',
        assignee: 'Jackson',
        platforms: ['instagram', 'linkedin'],
        history: []
    },
    {
        id: '2',
        date: '2025-11-22',
        time: '14:00',
        title: 'Dia do Engenheiro Eletricista',
        description: 'Homenagem aos profissionais parceiros.',
        caption: 'Parabéns a todos os engenheiros que iluminam nosso caminho! 💡⚡',
        category: 'solis_voce',
        type: 'static',
        status: 'in_progress',
        assignee: 'Beatriz',
        platforms: ['instagram', 'facebook'],
        history: []
    }
];

export const MOCK_ACCOUNTS_PAYABLE: AccountPayable[] = [
    { id: '1', supplier: 'Google Brasil', description: 'Google Ads - Nov/25', amount: 4500.00, dueDate: '2025-11-25', nfArrived: true, boletoArrived: true, status: 'pending', recurrence: 'monthly' },
    { id: '2', supplier: 'Meta Platforms', description: 'Facebook Ads - Nov/25', amount: 3200.00, dueDate: '2025-11-28', nfArrived: true, boletoArrived: false, status: 'pending', recurrence: 'monthly' },
    { id: '3', supplier: 'Adobe Systems', description: 'Licença Creative Cloud', amount: 280.00, dueDate: '2025-11-30', nfArrived: false, boletoArrived: false, status: 'pending', recurrence: 'monthly' }
];

export const MARKETING_CHANNELS_DATA: MarketingChannelData = {
    'NOV': [
        { channel: 'Google Ads', visits: 15000, leads: 450, conversion: 3.0, investment: 5000, cpl: 11.11, roas: 4.5 },
        { channel: 'Meta Ads', visits: 12000, leads: 380, conversion: 3.1, investment: 4000, cpl: 10.52, roas: 5.2 },
        { channel: 'Orgânico', visits: 8000, leads: 200, conversion: 2.5, investment: 0, cpl: 0, roas: 0 },
    ]
};

export const ANNUAL_CHANNEL_DATA: ChannelPerformance[] = [
    { channel: 'Google Ads', visits: 150000, leads: 4500, conversion: 3.0, investment: 50000, cpl: 11.11, roas: 4.5 },
    { channel: 'Meta Ads', visits: 120000, leads: 3800, conversion: 3.1, investment: 40000, cpl: 10.52, roas: 5.2 },
    { channel: 'Orgânico', visits: 80000, leads: 2000, conversion: 2.5, investment: 0, cpl: 0, roas: 0 },
];

export const MOCK_PDV_POSTS: PdvPost[] = [
    { id: '1', representative_uuid: 'uuid-do-andre', pdv_name: 'Elétrica Silva', post_date: '2025-11-15', month: 'NOV', platform: 'instagram', link: 'https://instagram.com/p/123', status: 'approved', created_at: '2025-11-15T10:00:00Z', updated_at: '2025-11-15T10:00:00Z' },
    { id: '2', representative_uuid: 'uuid-do-cesar', pdv_name: 'Solar Tech', post_date: '2025-11-18', month: 'NOV', platform: 'facebook', status: 'pending', created_at: '2025-11-18T14:00:00Z', updated_at: '2025-11-18T14:00:00Z' },
];

export const RECURRENT_PDVS: RecurrentPdv[] = [
    { id: '1', name: 'Elétrica Silva', representative_uuid: 'uuid-do-andre', city: 'Recife - PE', created_at: '2025-01-01T00:00:00Z', updated_at: '2025-01-01T00:00:00Z' },
    { id: '2', name: 'Solar Tech', representative_uuid: 'uuid-do-cesar', city: 'São Paulo - SP', created_at: '2025-01-01T00:00:00Z', updated_at: '2025-01-01T00:00:00Z' },
];

export const MOCK_SHOWROOM_ITEMS: ShowroomItem[] = [
    { id: '1', pdv: 'Elétrica Silva', city: 'Recife - PE', contact: 'João Silva', representative_uuid: 'uuid-do-andre', notDelivered: true, delivered: false, deliveryForecast: '2025-12-10', workshopDate: '' },
    { id: '2', pdv: 'Solar Tech', city: 'São Paulo - SP', contact: 'Maria Souza', representative_uuid: 'uuid-do-cesar', notDelivered: false, delivered: true, deliveryForecast: '', workshopDate: '2025-11-20' },
];

export const OFFLINE_ACTIONS_DATA: OfflineAction[] = [
    { solicitado: 'R$ 1.500,00', data: '10/11/2025', aprovado: 'R$ 1.500,00', pedido: '12345', saida: '12/11/2025', previsao: '15/11/2025', entrega: '15/11/2025', cidade: 'Recife', uf: 'PE', pontuado: 'SIM', status: 'Entregue', pdv: 'Elétrica Silva', representative_uuid: 'uuid-do-andre', category: 'AÇÃO COOPERADA' },
    { solicitado: 'R$ 800,00', data: '18/11/2025', aprovado: '', pedido: '', saida: '', previsao: '', entrega: '', cidade: 'São Paulo', uf: 'SP', pontuado: 'AINDA NÃO', status: 'Em análise', pdv: 'Solar Tech', representative_uuid: 'uuid-do-cesar', category: 'BRINDES' },
];

export const GIFT_ITEMS: GiftItem[] = [
    { name: 'Boné Solis', stock: 50, price: 25.00 },
    { name: 'Caneta Premium', stock: 100, price: 5.50 },
    { name: 'Caderno de Anotações', stock: 30, price: 15.00 },
    { name: 'Garrafa Térmica', stock: 15, price: 45.00 },
];

export const GIFT_TRANSACTIONS_MOCK: GiftTransaction[] = [
    { id: '1', itemName: 'Boné Solis', date: '01/11/2025', time: '09:00', quantity: 50, unit: 'unid.', type: 'in', price: 25.00 },
    { id: '2', itemName: 'Boné Solis', date: '05/11/2025', time: '14:00', quantity: 5, unit: 'unid.', type: 'out', representative: 'André' },
];

export const PROGRAM_CREDENTIALS: ProgramCredential[] = [
    { name: 'Instagram', user: 'solissolaroficial', password: '***', access: 'App Mobile' },
    { name: 'LinkedIn', user: 'solis-solar', password: '***', access: 'Web' },
    { name: 'RD Station', user: 'marketing@solis.ind.br', password: '***', access: 'Web' },
];

export const INTERNAL_CONTACTS: InternalContact[] = [
    { name: 'Suporte TI', role: 'Técnico', department: 'TI', extension: '100', email: 'suporte@solis.ind.br' },
    { name: 'RH', role: 'Analista', department: 'Recursos Humanos', extension: '101', email: 'rh@solis.ind.br' },
];