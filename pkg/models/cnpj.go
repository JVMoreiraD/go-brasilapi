package models

// Estrutura do objeto do cnpj
type CNPJ struct {
	Cnpj                       string  `json:"cnpj"`
	IdentificadorMatrizFilial  int     `json:"identificador_matriz_filial"`
	DescricaoMatrizFilial      string  `json:"descricao_matriz_filial"`
	RazaoSocial                string  `json:"razao_social"`
	NomeFantasia               string  `json:"nome_fantasia"`
	SituacaoCadastral          int     `json:"situacao_cadastral"`
	DescricaoSituacaoCadastral string  `json:"descricao_situacao_cadastral"`
	DataSituacaoCadastral      string  `json:"data_situacao_cadastral"`
	MotivoSituacaoCadastral    int     `json:"motivo_situacao_cadastral"`
	NomeCidadeExterior         string  `json:"nome_cidade_exterior"`
	CodigoNaturezaJuridica     int     `json:"codigo_natureza_juridica"`
	DataInicioAtividade        string  `json:"data_inicio_atividade"`
	CnaeFiscal                 int     `json:"cnae_fiscal"`
	CnaeFiscalDescricao        string  `json:"cnae_fiscal_descricao"`
	DescricaoTipoDeLogradouro  string  `json:"descricao_tipo_de_logradouro"`
	Logradouro                 string  `json:"logradouro"`
	Numero                     string  `json:"numero"`
	Complemento                string  `json:"complemento"`
	Bairro                     string  `json:"bairro"`
	Cep                        string  `json:"cep"`
	Uf                         string  `json:"uf"`
	CodigoMunicipio            int     `json:"codigo_municipio"`
	Municipio                  string  `json:"municipio"`
	DDDTelefone1               string  `json:"ddd_telefone_1"`
	DDDTelefone2               string  `json:"ddd_telefone_2"`
	DDDFax                     string  `json:"ddd_fax"`
	QualificacaoDoResponsavel  int     `json:"qualificacao_do_responsavel"`
	CapitalSocial              int     `json:"capital_social"`
	Porte                      string  `json:"porte"`
	DescricaoPorte             string  `json:"descricao_porte"`
	OpcaoPeloSimples           bool    `json:"opcao_pelo_simples"`
	DataOpcaoPeloSimples       string  `json:"data_opcao_pelo_simples"`
	DataExclusaoDoSimples      string  `json:"data_exclusao_do_simples"`
	OpcaoPeloMei               bool    `json:"opcao_pelo_mei"`
	SituacaoEspecial           string  `json:"situacao_especial"`
	DataSituacaoEspecial       string  `json:"data_situacao_especial"`
	CnaesSecundarios           []cnaes `json:"cnaes_secundarios"`
	Qsa                        []qsa   `json:"qsa"`
}

// Estrutura do objeto CNAES(Classificação Nacional das Atividades Econômicas.)
type cnaes struct {
	Codigo    int    `json:"codigo"`
	Descricao string `json:"descricao"`
}

// Este objeto do QSA(Quadro de sócios e administradores)
type qsa struct {
	IdentificadorDeSocio                 int     `json:"identificador_de_socio"`
	NomeSocio                            string  `json:"nome_socio"`
	CnpjCpfDoSocio                       string  `json:"cnpj_cpf_do_socio"`
	CodigoQualificacaoSocio              int     `json:"codigo_qualificacao_socio"`
	PercentualCapitalSocial              int     `json:"percentual_capital_social"`
	DataEntradaSociedade                 string  `json:"ata_entrada_sociedade"`
	CpfRepresentanteLegal                *string `json:"cpf_representante_legal"`
	NomeRepresentanteLegal               *string `json:"nome_representante_legal"`
	CodigoQualificacaoRepresentanteLegal *string `json:"codigo_qualificacao_representante_legal"`
}
