package env

import (
	"os"

	"github.com/PMoneda/flow"
	"github.com/mundipagg/boleto-api/config"
	"github.com/mundipagg/boleto-api/metrics"
	"github.com/mundipagg/boleto-api/models"
	"github.com/mundipagg/boleto-api/util"
)

func Config(devMode, mockMode, disableLog bool) {
	configFlags(devMode, mockMode, disableLog)
	flow.RegisterConnector("logseq", util.SeqLogConector)
	flow.RegisterConnector("apierro", models.BoletoErrorConector)
	flow.RegisterConnector("tls", util.TlsConector)
	metrics.Install()
}

func ConfigMock(port string) {
	os.Setenv("URL_BB_REGISTER_BOLETO", "http://localhost:"+port+"/registrarBoleto")
	os.Setenv("URL_BB_TOKEN", "http://localhost:"+port+"/oauth/token")
	os.Setenv("URL_CAIXA", "http://localhost:"+port+"/caixa/registrarBoleto")
	os.Setenv("URL_CITI", "http://localhost:"+port+"/citi/registrarBoleto")
	os.Setenv("URL_SANTANDER_TICKET", "tls://localhost:"+port+"/santander/get-ticket")
	os.Setenv("URL_SANTANDER_REGISTER", "tls://localhost:"+port+"/santander/register")
	os.Setenv("URL_BRADESCO_SHOPFACIL", "http://localhost:"+port+"/bradescoshopfacil/registrarBoleto")
	os.Setenv("URL_ITAU_TICKET", "http://localhost:"+port+"/itau/gerarToken")
	os.Setenv("URL_ITAU_REGISTER", "http://localhost:"+port+"/itau/registrarBoleto")
	os.Setenv("URL_BRADESCO_NET_EMPRESA", "http://localhost:"+port+"/bradesconetempresa/registrarBoleto")
	os.Setenv("MONGODB_URL", "localhost:27017")
	os.Setenv("MONGODB_USER", "")
	os.Setenv("MONGODB_PASSWORD", "")
	config.Install(true, true, true)
}

func configFlags(devMode, mockMode, disableLog bool) {
	if devMode {
		os.Setenv("INFLUXDB_HOST", "http://localhost")
		os.Setenv("INFLUXDB_PORT", "8086")
		os.Setenv("PDF_API", "http://localhost:7070/topdf")
		os.Setenv("API_PORT", "3000")
		os.Setenv("API_VERSION", "0.0.1")
		os.Setenv("ENVIROMENT", "Development")
		os.Setenv("SEQ_URL", "http://localhost:5341")
		os.Setenv("SEQ_API_KEY", "4jZzTybZ9bUHtJiPdh6")
		os.Setenv("ENABLE_REQUEST_LOG", "false")
		os.Setenv("ENABLE_PRINT_REQUEST", "true")
		os.Setenv("URL_BB_REGISTER_BOLETO", "https://cobranca.homologa.bb.com.br:7101/registrarBoleto")
		os.Setenv("URL_BB_TOKEN", "https://oauth.hm.bb.com.br/oauth/token")
		os.Setenv("CAIXA_ENV", "SGCBS01D")
		os.Setenv("URL_CAIXA", "https://des.barramento.caixa.gov.br/sibar/ManutencaoCobrancaBancaria/Boleto/Externo")
		os.Setenv("URL_CITI", "https://citigroupsoauat.citigroup.com/comercioeletronico/registerboleto/RegisterBoletoSOAP")
		os.Setenv("URL_CITI_BOLETO", "https://ebillpayer.uat.brazil.citigroup.com/ebillpayer/jspInformaDadosConsulta.jsp")
		os.Setenv("APP_URL", "http://localhost:3000/boleto")
		os.Setenv("ELASTIC_URL", "http://localhost:9200")
		os.Setenv("MONGODB_URL", "localhost:27017")
		os.Setenv("MONGODB_USER", "")
		os.Setenv("MONGODB_PASSWORD", "")
		os.Setenv("BOLETO_JSON_STORE", "C:\\boleto_json_store\\JSONFileStore.json")
		os.Setenv("CERT_BOLETO_CRT", "C:\\cert_boleto_api\\certificate.crt")
		os.Setenv("CERT_BOLETO_KEY", "C:\\cert_boleto_api\\mundi.key")
		os.Setenv("CERT_BOLETO_CA", "C:\\cert_boleto_api\\ca-cert.ca")
		os.Setenv("CERT_ICP_BOLETO_KEY", "C:\\cert_boleto_api\\ICP_PKey.key")
		os.Setenv("CERT_ICP_BOLETO_CHAIN_CA", "C:\\cert_boleto_api\\ICP_cadeiaCerts.pem")
		os.Setenv("URL_SANTANDER_TICKET", "https://ymbdlb.santander.com.br/dl-ticket-services/TicketEndpointService")
		os.Setenv("URL_SANTANDER_REGISTER", "https://ymbcash.santander.com.br/ymbsrv/CobrancaEndpointService")
		os.Setenv("URL_BRADESCO_SHOPFACIL", "https://homolog.meiosdepagamentobradesco.com.br/apiboleto/transacao")
		os.Setenv("ITAU_ENV", "1")
		os.Setenv("SANTANDER_ENV", "T")
		os.Setenv("URL_ITAU_REGISTER", "https://gerador-boletos.itau.com.br/router-gateway-app/public/codigo_barras/registro")
		os.Setenv("URL_ITAU_TICKET", "https://oauth.itau.com.br/identity/connect/token")
		os.Setenv("URL_BRADESCO_NET_EMPRESA", "https://cobranca.bradesconetempresa.b.br/ibpjregistrotitulows/registrotitulohomologacao")
	}
	config.Install(mockMode, devMode, disableLog)
}
