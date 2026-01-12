#!/bin/bash
set -e

# Cores para output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Iniciando deploy${NC}"

# Carregar variáveis de ambiente do arquivo .env.ecr
if [ -f .env.ecr ]; then
    export $(cat .env.ecr | xargs)
else
    echo -e "${RED}❌ Arquivo .env.ecr não encontrado!${NC}"
    exit 1
fi

# Verificar se as variáveis foram carregadas
if [ -z "$AWS_ACCOUNT_ID" ] || [ -z "$AWS_REGION" ]; then
    echo -e "${RED}❌ Variáveis AWS_ACCOUNT_ID ou AWS_REGION não definidas no .env.ecr${NC}"
    exit 1
fi

echo -e "${YELLOW}📦 AWS Account: $AWS_ACCOUNT_ID${NC}"
echo -e "${YELLOW}🌎 AWS Region: $AWS_REGION${NC}"

# Fazer login no ECR
echo -e "${GREEN}🔐 Fazendo login no ECR...${NC}"
aws ecr get-login-password --region $AWS_REGION | docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com

# Baixar as imagens mais recentes com a tag :prod
echo -e "${GREEN}📥 Fazendo pull das imagens...${NC}"

echo "  → Backend..."
docker pull $AWS_ACCOUNT_ID.dkr.ecr.$AWS_REGION.amazonaws.com/solis:prod

# Reiniciar os containers com as novas imagens
echo -e "${GREEN}🔄 Restart dos containers...${NC}"
docker-compose -f docker-compose.yml up -d --no-build

echo -e "${GREEN}✅ Deploy completo!${NC}"
echo ""
echo -e "${YELLOW}📊 Status dos containers:${NC}"
docker-compose -f docker-compose.yml ps

echo ""
echo -e "${GREEN}🎉 Deploy finalizado com sucesso!${NC}"