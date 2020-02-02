package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// Company contains information about a single company
type Company struct {
	ID            int    `json:"id" binding:"required"`
	Likes         int    `json:"likes"`
	Company       string `json:"company" binding:"required"`
	FoundedYear   int    `json:"founded_year" binding:"required"`
	Description   string `json:"description" binding:"required"`
	OfficialURL   string `json:"official_url" binding:"required"`
	CrunchbaseURL string `json:"crunchbase_url" binding:"required"`
}

var companies = []Company{
	Company{1, 0, "120WaterAudit", 2016, "120WaterAudit offers cloud-based water management software.", "https://120wateraudit.com/", "https://www.crunchbase.com/organization/120wateraudit#section-overview"},
	Company{2, 0, "3AM Innovations", 2015, "3AM has designed an IPS solution that utilizes state-of-the-art sensors, communication technologies, and software called FLARE.", "https://www.3aminnovations.com/", "https://www.crunchbase.com/organization/3am-innovations"},
	Company{3, 0, "Accela", 1999, "Accela provides civic engagement solutions and improves core processes for city, county, state and federal governments.", "https://www.accela.com/", "https://www.crunchbase.com/organization/accela#section-overview"},
	Company{4, 0, "Acivilate", 2014, "They are a Software-as-a-Service (SaaS) company committed to reducing recidivism.", "https://acivilate.com/", "https://www.crunchbase.com/organization/acivilate"},
	Company{5, 0, "APPCityLife®, Inc.", 2009, "APPCityLife is an end-to-end platform transforming the way cities and agencies produce mobile apps.", "https://www.appcitylife.com/", "https://www.crunchbase.com/organization/appcitylife"},
	Company{6, 0, "ArchiveSocial", 2011, "ArchiveSocial is a social media archiving solution for business compliance and legal protection activities in enterprises.", "https://archivesocial.com/", "https://www.crunchbase.com/organization/archivesocial#section-overview"},
	Company{7, 0, "Avenu Insights & Analytics", 1989, "Avenu Insights & Analytics is the provider of revenue enhancement, technology, and consulting services for government agencies.", "https://www.avenuinsights.com/", "https://www.crunchbase.com/organization/avenu-insights-analytics#section-overview"},
	Company{8, 0, "Axon", 1993, "Axon provides electronic control devices to law enforcement and corrections agencies.", "https://www.appcitylife.com/", "https://www.crunchbase.com/organization/appcitylife"},
	Company{9, 0, "Balancing Act by Engaged Public", 1998, "Balancing Act is a suite of tools to help government engage citizens on budget priorities and financial issues. Product of Engaged Public.", "https://www.axon.com/", "https://www.crunchbase.com/organization/taser-international"},
	Company{10, 0, "Bang the Table", 2007, "Bang the Table offers an online citizen engagement platform for local government.", "https://bangthetable.com/", "https://www.crunchbase.com/organization/bang-the-table"},
	Company{11, 0, "Binti", 2014, "Binti is a for-profit mission driven tech company with the goal of helping children have safe, loving and stable families.", "https://binti.com/", "https://www.crunchbase.com/organization/binti"},
	Company{12, 0, "Biobot Analytics", 2017, "Biobot Analytics analyzes city sewage to estimate opioid consumption.", "http://www.biobot.io/", "https://www.crunchbase.com/organization/biobot-analytics"},
	Company{13, 0, "BlueLine Grid", 2013, "BlueLine Grid operates the premier mass notification and critical collaboration platform for corporations and government agencies.", "https://www.bluelinegrid.com/", "https://www.crunchbase.com/organization/bratton-technologies-inc"},
	Company{14, 0, "BondLink", 2016, "BondLink builds technology products and tools to help modernize the municipal bond market.", "https://www.bondlink.com/", "https://www.crunchbase.com/organization/bondlink"},
	Company{15, 0, "Callyo", 2010, "Advancing mobile technology for law enforcement investigators to create a safer world for children", "https://callyo.com/", "https://www.crunchbase.com/organization/callyo"},
	Company{16, 0, "Calytera", 1989, "CSDC provides software to help governments efficiently manage resources & processes to enable growth and seamlessly interact with citizens.", "https://www.calytera.com/", "https://www.crunchbase.com/organization/csdc-systems"},
	Company{17, 0, "Cardinality.ai", 2017, "Super Charged Citizen Experience, Out of the Box!", "http://www.cardinality.ai/", "https://www.crunchbase.com/organization/cardinality-ai"},
	Company{18, 0, "Cartegraph", 1994, "Cartegraph makes modern operations management software for local governments.", "https://www.cartegraph.com/", "https://www.crunchbase.com/organization/cartegraph"},
	Company{19, 0, "Casebook PBC", 2017, "We are the developer of the patent-pending, award-winning Casebook SaaS platform for Human Services.", "http://www.casebook.net/", "https://www.crunchbase.com/organization/casebook-pbc"},
	Company{20, 0, "CentralSquare Technologies", 1979, "CentralSquare is a new type of technology leader innovating on behalf of the public sector.", "https://www.centralsqr.com/", "https://www.crunchbase.com/organization/centralsquare-technologies"},
	Company{21, 0, "Citibot", 2016, "Citibot is the tool for citizens and their governments to use for communication and civic change.", "https://www.citibot.io/", "https://www.crunchbase.com/organization/citibot"},
	Company{22, 0, "Citymart", 2011, "Citymart transforms the way cities solve problems, connecting them with new ideas through open challenges to entrepreneurs and citizens.", "http://citymart.com/", "https://www.crunchbase.com/organization/citymart"},
	Company{23, 0, "CivicActions", 2004, "CivicActions helps agencies and organizations successfully execute large-scale software projects using open standards technologies.", "https://civicactions.com/", "https://www.crunchbase.com/organization/civicactions"},
	Company{24, 0, "CivicPlus", 1994, "CivicPlus provides integrated software for cities and counties across North America. (Websites, HRMS, Mass Notfications and more)", "http://www.civicplus.com/", "https://www.crunchbase.com/organization/civic-plus#section-overview"},
	Company{25, 0, "Civis Analytics", 2013, "Civis Analytics helps leading commercial, nonprofit and government organizations leverage data to make better decisions.", "http://www.civisanalytics.com/", "https://www.crunchbase.com/organization/civis-analytics"},
	Company{26, 0, "Clear Ballot Group, Inc.", 2009, "Clear Ballot is re-imagining elections with the world's most innovative and transparent voting system.", "http://www.clearballot.com/", "https://www.crunchbase.com/organization/clear-ballot-group-inc"},
	Company{27, 0, "ClearGov", 2015, "ClearGov is helping citizens better understand their local government's financial performance.", "https://www.cleargov.com/", "https://www.crunchbase.com/organization/cleargov"},
	Company{28, 0, "CoInspect", 2014, "CoInspect builds mobile software to manage compliance, quality assurance, and brand standards.", "https://coinspectapp.com/", "https://www.crunchbase.com/organization/mewe-inc-"},
	Company{29, 0, "Compology", 2012, "Compology offers camera monitoring for waste container management.", "http://www.compology.com/", "https://www.crunchbase.com/organization/compology"},
	Company{30, 0, "Computronix", 1979, "Computronix provides permitting, licensing, inspection and compliance solutions for government, based on our award-winning POSSE platform.", "https://www.computronix.com/", "https://www.crunchbase.com/organization/computronix"},
	Company{31, 0, "Coord", 2016, "Coord is a mobility company that creates seamless mobility and self-driving experiences today through deep integrations.", "https://www.coord.com/", "https://www.crunchbase.com/organization/coord#section-overview"},
	Company{32, 0, "coUrbanize", 2013, "coUrbanize is an online community engagement platform connecting real estate developers and municipalities with residents.", "https://www.courbanize.com/", "https://www.crunchbase.com/organization/courbanize"},
	Company{33, 0, "Cubic Corporation", 1951, "Cubic Corporation (Cubic) is an international provider of systems and solutions that address the mass transit and global defense markets.", "http://cubic.com/", "https://www.crunchbase.com/organization/cubic-corporation"},
	Company{34, 0, "EasyVote Solutions", 2013, "EasyVote Solutions delivers a SaaS platform to city, county and state election offices to help manage the process of running elections.", "http://www.easyvotesolutions.com/", "https://www.crunchbase.com/organization/easyvote-solutions"},
	Company{35, 0, "Electro Scan Inc.", 2011, "Designs, develops, and markets, next generation sensors and cloud applications to assess water, sewer, and gas pipelines.", "https://www.electroscan.com/", "https://www.crunchbase.com/organization/electro-scan-inc"},
	Company{36, 0, "Elucd", 2016, "Elucd is a GovTech company that provides cities with near-realtime understanding of how citizens feel about their government.", "https://www.elucd.com/", "https://www.crunchbase.com/organization/elucd"},
	Company{37, 0, "Esri", 1969, "Esri offers geographic information system software, web GIS, and geodatabase management applications.", "https://www.esri.com/en-us/home", "https://www.crunchbase.com/organization/esri"},
	Company{38, 0, "FATHOM", 2003, "FATHOM are transforming the water utility paradigm by unleashing the power of data.", "http://www.gwfathom.com/", "https://www.crunchbase.com/organization/fathom-3"},
	Company{39, 0, "Forensic Logic", 2003, "Forensic Logic is a company founded on a simple idea: there is too much crime on our streets because critical information is inaccessible", "http://www.forensiclogic.com/", "https://www.crunchbase.com/organization/forensic-logic"},
	Company{40, 0, "GCR Inc.", 1979, "GCR Inc., a professional services firm, delivers expert consulting services and cutting-edge technology solutions", "http://www.gcrinc.com/", "https://www.crunchbase.com/organization/gcr-associates"},
	Company{41, 0, "GovPilot", 2014, "GovPilot is a web-based Management Platform developed exclusively for local government.", "https://www.govpilot.com/", "https://www.crunchbase.com/organization/govpilot"},
	Company{42, 0, "GovQA", 2002, "GovQA, provides SaaS-based, collaboration tools enabling the nation's top cities, counties, and state agencies process complex work.", "http://www.govqa.com/", "https://www.crunchbase.com/organization/govqa"},
	Company{43, 0, "GovSense", 2014, "GovSense is cloud-based permitting, licensing and financial software for state and local government.", "http://www.govsense.com/", "https://www.crunchbase.com/organization/govsense"},
	Company{44, 0, "Granicus", 1999, "Granicus provides technology that empowers government organizations to create better lives for the people they serve.", "https://granicus.com/", "https://www.crunchbase.com/organization/granicus"},
	Company{45, 0, "Gridics", 2015, "Intelligent 3D Zoning and Data Platform", "https://www.gridics.com/", "https://www.crunchbase.com/organization/gridics"},
	Company{46, 0, "GTY Technology Holdings", 2016, "GTY Holdings is a gov tech acquisitions company comprising a number of smaller startups.", "https://www.gtytechnology.com/", "https://www.crunchbase.com/organization/gty-technology-holdings"},
	Company{47, 0, "HAAS Alert", 2015, "HAAS uses mobile data to alert drivers (and cyclists) of approaching emergency vehicles through vehicle-to-vehicle notifications.", "http://www.haasalert.com/", "https://www.crunchbase.com/organization/haasalert"},
	Company{48, 0, "IPS Group", 1995, "IPS Group globally delivers smart city tech within an Internet of Things framework.", "http://ipsgroupinc.com/", "https://www.crunchbase.com/organization/ips-group"},
	Company{49, 0, "Itron", 1977, "Itron is a technology company that offers products and services on energy and water resource management.", "https://www.itron.com/na/", "https://www.crunchbase.com/organization/itron"},
	Company{50, 0, "LiveStories", 2013, "LiveStories is a data hub that includes a civic data library, team collaboration, and online publishing—all on one platform.", "https://www.livestories.com/", "https://www.crunchbase.com/organization/livestories"},
	Company{51, 0, "LotaData, Inc.", 2015, "LOTADATA transforms time, place, activity into actionable context, insights and segments for mobile analytics and business intelligence", "https://lotadata.com/", "https://www.crunchbase.com/organization/lotadata--inc-"},
	Company{52, 0, "Mark43", 2012, "Mark43 develops law enforcement software that allows police to collect, manage, analyze and share information.", "http://mark43.com/", "https://www.crunchbase.com/organization/mark43"},
	Company{53, 0, "Maximus", 1975, "Maximus is an operator of government health and human services programs that help serve citizens with compassion and efficacy.", "http://www.maximus.com/", "https://www.crunchbase.com/organization/maximus"},
	Company{54, 0, "Moovit", 2012, "Moovit is leading Mobility as a Service (MaaS) provider and #1 public transit app.", "https://moovit.com/", "https://www.crunchbase.com/organization/moovitapp"},
	Company{55, 0, "Motorola Solutions", 1928, "Motorola Solutions creates mission-critical communication solutions and services for public safety and commercial customers.", "http://www.motorolasolutions.com/", "https://www.crunchbase.com/organization/motorola-solutions"},
	Company{56, 0, "Munetrix", 2010, "Munetrix is a web-based information source providing financial information and forecasting tools for municipal governments and schools.", "http://munetrix.com/", "https://www.crunchbase.com/organization/munetrix"},
	Company{57, 0, "Municode", 1951, "Municode provides website design & development, electronic payment processing, codification, and publishing services to local governments.", "https://www.municode.com/", "https://www.crunchbase.com/organization/municipal-code-corporation"},
	Company{58, 0, "Neighborland", 2011, "Neighborland enables residents to collaborate with local organizations and take action on their area-related issues.", "http://neighborland.com/", "https://www.crunchbase.com/organization/neighborland"},
	Company{59, 0, "NEOGOV", 1999, "NEOGOV is the market and technology leader in on-demand human resources software for the public sector.", "http://www.neogov.com/", "https://www.crunchbase.com/organization/neogov"},
	Company{60, 0, "Nextdoor", 2010, "Nextdoor is the world's largest private social network used by 90% of US neighborhoods and over 5,000 public agencies to build stronger, safer, happier communities.", "http://nextdoor.com/", "https://www.crunchbase.com/organization/nextdoor"},
	Company{61, 0, "NextRequest", 2015, "NextRequest makes public records requests friendlier for the public and easier for governments.", "https://www.nextrequest.com/", "https://www.crunchbase.com/organization/nextrequest"},
	Company{62, 0, "NIC", 1992, "NIC is a provider of innovative digital government solutions and secure payment processing.", "http://www.egov.com/", "https://www.crunchbase.com/organization/nic"},
	Company{63, 0, "Numetric", 2015, "Numetric works with Departments of Transportation in making their roadways better and safer, through cloud-based software and data services.", "http://www.numetric.com/", "https://www.crunchbase.com/organization/numetric"},
	Company{64, 0, "One Concern", 2015, "One Concern is a disaster solutions company that provides damage estimates using artificial intelligence on natural phenomena sciences.", "https://www.oneconcern.com/", "https://www.crunchbase.com/organization/oneconcern"},
	Company{65, 0, "OpenDataSoft", 2011, "OpenDataSoft is a cloud-based turnkey platform for data publishing and API management.", "http://www.opendatasoft.com/", "https://www.crunchbase.com/organization/opendatasoft"},
	Company{66, 0, "OpenGov", 2012, "OpenGov is the first complete cloud solution for public sector budgeting, operational performance, and citizen engagement.", "https://opengov.com/", "https://www.crunchbase.com/organization/opengov"},
	Company{67, 0, "Passport", 2010, "Passport specializes in enterprise business applications and payments for parking and transportation.", "https://passportshipping.com/", "https://www.crunchbase.com/organization/passport"},
	Company{68, 0, "PayIt", 2013, "PayIt simplifies doing business with state, local and federal government through it's mobile transaction and payment platform.", "https://www.payitgov.com/", "https://www.crunchbase.com/organization/payit"},
	Company{69, 0, "Periscope Holdings", 2001, "Periscope Holdings helps organizations build a better procurement process.", "http://www.periscopeholdings.com/", "https://www.crunchbase.com/organization/periscope-holdings"},
	Company{70, 0, "Planet", 2010, "Planet builds small satellites and delivers information about the changing planet.", "https://www.planet.com/", "https://www.crunchbase.com/organization/planet-labs"},
	Company{71, 0, "Pondera Solutions", 2011, "Pondera offers a comprehensive cloud solution to help you detect, investigate, and enforce fraud, waste, and abuse.", "http://www.ponderasolutions.com/", "https://www.crunchbase.com/organization/pondera-solutions"},
	Company{72, 0, "PredPol", 2012, "PredPol predicts crime using cloud software technology that identifies the highest risk times and places in near real-time.", "http://www.predpol.com/", "https://www.crunchbase.com/organization/predpol"},
	Company{73, 0, "PrimeGov", 2014, "PrimeGov is committed to providing world-class customer service in its delivery of innovative legislative management solutions.", "http://primegov.com/", "https://www.crunchbase.com/organization/primegov"},
	Company{74, 0, "Promise", 0, "Promise provides a cost-effective, humane alternative to incarceration by improving long-term outcomes for individuals and communities.", "https://joinpromise.com/", "https://www.crunchbase.com/organization/promise"},
	Company{75, 0, "ProudCity", 2016, "ProudCity is a software company providing cities with websites and online government services.", "https://proudcity.com/", "https://www.crunchbase.com/organization/proudcity"},
	Company{76, 0, "Quicket Solutions", 2013, "Quicket provides a cloud-based data management and operational intelligence platform for sensitive government workloads.", "https://www.quicketsolutions.com/", "https://www.crunchbase.com/organization/quicket-solutions"},
	Company{77, 0, "Rachio", 2012, "Rachio is a technology company offering a smart irrigation controller with cloud-based software and web-based dashboards.", "http://rachio.com/", "https://www.crunchbase.com/organization/rachio"},
	Company{78, 0, "RapidDeploy", 2013, "RapidDeploy is a Cloud–based Computer Aided Dispatch (CAD) and Analytics platform.", "http://www.rapiddeploy.com/", "https://www.crunchbase.com/organization/rapiddeploy"},
	Company{79, 0, "RapidSOS", 2013, "RapidSOS is an advanced emergency communication company.", "https://rapidsos.com/", "https://www.crunchbase.com/organization/rapidsos"},
	Company{80, 0, "Remix", 2014, "Remix is a planning platform for public transit.", "https://www.remix.com/", "https://www.crunchbase.com/organization/remix"},
	Company{81, 0, "Replica", 0, "Replica is a next-generation urban planning tool that can help cities answer key transportation questions", "https://replicahq.com/", "https://www.crunchbase.com/organization/replica#section-overview"},
	Company{82, 0, "RoadBotics", 2016, "AI technology that monitors and manages roadways around the world.", "https://www.roadbotics.com/", "https://www.crunchbase.com/organization/roadbotics-inc"},
	Company{83, 0, "Rock Solid Technologies", 1994, "Rock Solid Technologies is a software products research and development company.", "https://www.rocksolid.com/", "https://www.crunchbase.com/organization/rock-solid-technologies"},
	Company{84, 0, "Sagitec Solutions LLC", 2004, "We design & deliver tailor-made pension, provident fund, unemployment insurance, and healthcare and life sciences software solutions.", "https://www.sagitec.com/", "https://www.crunchbase.com/organization/sagitec-solutions-llc"},
	Company{85, 0, "SeamlessDocs", 2011, "SeamlessDocs is a eSig and form automation platform that specializes in working with govs to help them digitize their PDF & form processes.", "http://www.seamlessdocs.com/", "https://www.crunchbase.com/organization/seamlessdocs"},
	Company{86, 0, "ShotSpotter", 1996, "ShotSpotter, a public company, is the global leader in gunfire detection and location technology. It is used by more than 90 cities.", "https://www.shotspotter.com/", "https://www.crunchbase.com/organization/shotspotter"},
	Company{87, 0, "Sidewalk Labs", 2014, "Sidewalk Labs works with cities to build products that address urban problems.", "http://www.sidewalkinc.com/", "https://www.crunchbase.com/organization/sidewalk-labs"},
	Company{88, 0, "Smarking", 2014, "Smarking turns parking data into instant actionable insights.", "https://smarking.com/", "https://www.crunchbase.com/organization/smarking"},
	Company{89, 0, "StreetLight Data", 2010, "StreetLight Data is a data analytics company transforming urban planning and transportation design with the power of geospatial data.", "http://www.streetlightdata.com/", "https://www.crunchbase.com/organization/streetlight-data"},
	Company{90, 0, "Swiftly", 2014, "Swiftly is changing transportation, mobility, and cities through better data", "http://goswift.ly/", "https://www.crunchbase.com/organization/swiftlyinc"},
	Company{91, 0, "The Boring Company", 2016, "The Boring Company aims to dig tunnels efficiently to facilitate an underground transportation network.", "https://boringcompany.com/", "https://www.crunchbase.com/organization/the-boring-company"},
	Company{92, 0, "Tyler Technologies", 1966, "Tyler Technologies is a software company providing integrated software and technology services to the public sector.", "http://www.tylertech.com/", "https://www.crunchbase.com/organization/tyler-technologies"},
	Company{93, 0, "UrbanLeap", 2017, "UrbanLeap offers software that helps governments run pilot projects to test new technologies.", "http://www.urbanleap.io/", "https://www.crunchbase.com/organization/urbanleap-co"},
	Company{94, 0, "Utilidata", 1983, "Utilidata works with utilities to save energy, increase reliability and better detect grid anomalies.", "http://www.utilidata.com/", "https://www.crunchbase.com/organization/utilidata"},
	Company{95, 0, "Utilis Inc.", 2013, "Utilis uses satellite imagery to monitor underground water systems and detect leaks.", "http://utiliscorp.com/", "https://www.crunchbase.com/organization/utilis-israel-inc-co"},
	Company{96, 0, "Varuna", 2018, "Varuna provides alerts, recommendations, and predictions to help water utility leaders to effectively manage a water system.", "https://www.varunaiot.com/", "https://www.crunchbase.com/organization/varuna-co"},
	Company{97, 0, "Visionary Integration Professionals", 1996, "VIP makes business strategy software for governments and corporations.", "http://www.trustvip.com/", "https://www.crunchbase.com/organization/visionary-integration-professionals"},
	Company{98, 0, "Voyage", 2017, "Voyage serves communities with autonomous vehicles", "http://voyage.auto/", "https://www.crunchbase.com/organization/voyage"},
	Company{99, 0, "WaterSmart Software", 2009, "WaterSmart Software uses mobile and online tools to help water utilities educate and engage their customers to save water and money.", "http://www.watersmart.com/", "https://www.crunchbase.com/organization/watersmart-software"},
	Company{100, 0, "Waycare", 2016, "WayCare optimizes traffic management systems leveraging predictive analytics, and enables two way vehicle to city communication.", "http://www.waycaretech.com/", "https://www.crunchbase.com/organization/waycare"},
	Company{101, 0, "AmigoCloud", 2011, "AmigoCloud is a geospatial platform that helps you collect, manage, analyze, visualize, and publish your location data.", "https://www.amigocloud.com/en/", "https://www.crunchbase.com/organization/amigocloud"},
	Company{102, 0, "Boundless Spatial", 2012, "Boundless is a provider of open source product support, training and core development to meet geospatial requirements.", "https://www.boundless.com/", "https://www.crunchbase.com/organization/boundless"},
	Company{103, 0, "Buildingeye", 2012, "Buildingeye is a user friendly planning information portal specializing in GIS and data visualization.", "https://buildingeye.com/", "https://www.crunchbase.com/organization/buildingeye"},
	Company{104, 0, "CityInsight", 0, "CityInsight creates mobile and desktop applications for municipal government.", "https://one.rocksolid.com/", "https://www.crunchbase.com/organization/cityinsight"},
	Company{105, 0, "CitySourced", 2006, "CitySourced is an enterprise civic engagement platform in the world, connecting citizens with their local government.", "https://www.citysourced.com/", "https://www.crunchbase.com/organization/citysourced"},
	Company{106, 0, "Cityzenith", 2009, "Cityzenith provides the world's leading data visualization platform software and applications for IoT and Smart Cities.", "http://www.cityzenith.com/", "https://www.crunchbase.com/organization/cityzenith"},
	Company{107, 0, "CivicScape", 2017, "A New Standard for Real-Time Policing", "https://www.civicscape.com/", "https://www.crunchbase.com/organization/civicscape"},
	Company{108, 0, "CivicSmart", 2015, "CivicSmart is a technology services and engineering company specializing in Smart City parking solutions.", "http://www.civicsmart.com/", "https://www.crunchbase.com/organization/civicsmart-inc"},
	Company{109, 0, "CIVIQ Smartscapes", 2015, "CIVIQ Smartscapes designs and manufactures interactive smart city communications structures", "http://www.civiqsmartscapes.com/", "https://www.crunchbase.com/organization/civiq-smartscapes"},
	Company{110, 0, "GRIDSMART Technologies, Inc.", 2006, "GRIDSMART uses a single camera to actuate intersections and gather data on arterials and highways.", "https://gridsmart.com/", "https://www.crunchbase.com/organization/aldis"},
	Company{111, 0, "Hyperloop Transportation Technologies", 2013, "Hyperloop Transportation Technologies is a global team that focuses on removing barriers of speed and congestion in travel.", "http://hyperloop.global/", "https://www.crunchbase.com/organization/hyperloop-transportation-technologies"},
	Company{112, 0, "Neighborly", 2012, "Neighborly modernizes access to public finance, the $1 billion dollar per day market that funds vital public projects like schools & parks.", "https://neighborly.com/", "https://www.crunchbase.com/organization/neighborly"},
	Company{113, 0, "SeeClickFix", 2008, "SeeClickFix provides tools for residents and governments to communicate for all sizes, populations, and budgets.", "http://www.seeclickfix.com/", "https://www.crunchbase.com/organization/seeclickfix"},
	Company{114, 0, "ViewPoint", 1995, "ViewPoint Cloud makes it easy to manage permit and license requests from intake to issuance.", "http://www.viewpointcs.com/", "https://www.crunchbase.com/organization/viewpoint"},
	Company{115, 0, "Votem Corp", 2014, "Votem enables citizens around the world to easily vote online with unprecedented verifiability, accessibility, security, and transparency.", "http://www.votem.com/", "https://www.crunchbase.com/organization/votem-corp"},
	Company{116, 0, "Aclara Technologies", 1972, "Aclara Technologies is a supplier of smart infrastructure solutions (SIS) to water, gas, and electric utilities globally.", "", "https://www.crunchbase.com/organization/aclara-technologies"},
	Company{117, 0, "CITYBASE", 2015, "CityBase is a gov tech platform that connects cities and citizens to make government more accessible and efficient.", "", "https://www.crunchbase.com/organization/409555065"},
	Company{118, 0, "Court Innovations", 2013, "Court Innovations develops and implements online negotiation systems for courts and constituents.", "", "https://www.crunchbase.com/organization/court-innovations"},
	Company{119, 0, "CRIMEWATCH", 2012, "Crowd Sourced Crime Fighting", "", "https://www.crunchbase.com/organization/crimewatch"},
	Company{120, 0, "Envisio Solutions Inc.", 2012, "Envisio helps governments achieve greater transparency & accountability by making it easy to track and report on strategic plans.", "", "https://www.crunchbase.com/organization/envisio-solutions-inc"},
	Company{121, 0, "Govlist", 2016, "Govlist is a procurement solution helping governments and vendors make better decisions, move faster, and simplify the procurement process.", "", "https://www.crunchbase.com/organization/govlist"},
	Company{122, 0, "OpenCounter", 2012, "OpenCounter combines Silicon Valley and public-sector talent to create best-of-class products for modern governments.", "", "https://www.crunchbase.com/organization/opencounter"},
	Company{123, 0, "Paladin Data Systems", 1994, "Paladin Data streamlines community development. SMARTGOV is SaaS solution that streamlines permitting, licensing, code enforcement and more.", "", "https://www.crunchbase.com/organization/paladin-data-systems"},
	Company{124, 0, "PermitZone", 2016, "Providing contractors and DIY pro’s, visibility into municipal knowledge, permitting requirements and the ability to pull permits online.", "", "https://www.crunchbase.com/organization/permitzone"},
	Company{125, 0, "Pluto AI", 2016, "Pluto is an analytics platform for smart water management.", "", "https://www.crunchbase.com/organization/pluto-ai"},
	Company{126, 0, "PublicInput.com", 2014, "PublicInput.com helps organizations engage their audience and collect valuable insights.", "", "https://www.crunchbase.com/organization/cityzen"},
	Company{127, 0, "Questica, Inc.", 1998, "Questica’s software suite includes fully featured web-based budgeting, performance and data visualization solutions.", "", "https://www.crunchbase.com/organization/questica-inc"},
	Company{128, 0, "Safe Fleet Holdings", 2013, "Leading global provider of safety solutions for fleet vehicles.", "", "https://www.crunchbase.com/organization/safe-fleet-holdings"},
	Company{129, 0, "SceneDoc", 2011, "SceneDoc is a provider of the leading smartphone/tablet-based investigation and field documentation platform.", "", "https://www.crunchbase.com/organization/-scenedoc"},
	Company{130, 0, "Seabourne", 2010, "Seabourne is a technology and data strategy firm.", "", "https://www.crunchbase.com/organization/seabourne"},
	Company{131, 0, "Securus Technologies", 1986, "Securus Technologies, Inc. is one of the largest providers of detainee communications, parolee tracking, and government information", "", "https://www.crunchbase.com/organization/securus-technologies"},
	Company{132, 0, "SmartProcure", 2011, "SmartProcure is a procurement database company for businesses and government agencies.", "", "https://www.crunchbase.com/organization/smartprocure"},
	Company{133, 0, "Socrata", 2007, "Socrata is a cloud-based software company, enables users to access the data of PS organizations via web, mobile, and M2M interfaces.", "", "https://www.crunchbase.com/organization/socrata"},
	Company{134, 0, "SPIDR Tech", 2015, "SPIDR Tech is a law enforcement technology development company.", "", "https://www.crunchbase.com/organization/spidr-tech"},
	Company{135, 0, "SpotCrime", 2007, "SpotCrime is an online crime data and mapping aggregation website.", "", "https://www.crunchbase.com/organization/spotcrime"},
	Company{136, 0, "TriTech Software Systems", 1992, "TriTech Software Systems develops highly integrated public safety products and services for police, fire, and EMS agencies worldwide.", "", "https://www.crunchbase.com/organization/tritech"},
	Company{137, 0, "Valor Water Analytics", 2013, "Valor Water Analytics develops customized financial data and dashboard tools for water utilities and businesses.", "", "https://www.crunchbase.com/organization/valor-water-analytics"},
	Company{138, 0, "VaultRMS", 2013, "VaultRMS is a cloud-based technology platform for public safety agencies.", "", "https://www.crunchbase.com/organization/vaultrms"},
	Company{139, 0, "Vision Internet", 1995, "Vision is a website design, development, hosting, and software company for local governments.", "", "https://www.crunchbase.com/organization/vision-internet"},
	Company{140, 0, "Vizalytics Technology", 2012, "Vizalytics creates location-specific business intelligence from open and private data for enterprise, government, and institutional clients.", "", "https://www.crunchbase.com/organization/468381163"},
	Company{141, 0, "Xcential Legislative Technologies", 2002, "Xcential provides SaaS-based solutions enable legislatures and regulatory to more efficiently create/publish legislation and regulations.", "", "https://www.crunchbase.com/organization/xcential-legislative-technologies"},
	Company{142, 0, "Appallicious", 2010, "Appallicious is an open data visualization company that creates products to help government better serve its citizens.", "", "https://www.crunchbase.com/organization/appallicious"},
	Company{143, 0, "Aunt Bertha", 2010, "Aunt Bertha enables individuals to find and apply for government and charitable social service programs.", "", "https://www.crunchbase.com/organization/aunt-bertha"},
	Company{144, 0, "Bidgely", 2010, "Bidgely is a disruptive technology company developing an energy monitoring and management solution for eco-friendly energy saving.", "", "https://www.crunchbase.com/organization/bidgely"},
	Company{145, 0, "Citizinvestor", 2012, "Citizinvestor is a crowdfunding and civic engagement platform for government projects.", "", "https://www.crunchbase.com/organization/citizinvestor"},
	Company{146, 0, "CityGrows", 2015, "CityGrows is the transparency-first process management platform for government.", "", "https://www.crunchbase.com/organization/city-grows"},
	Company{147, 0, "CityView", 1982, "CityView is a developer of local government business process automation software.", "", "https://www.crunchbase.com/organization/cityview"},
	Company{148, 0, "Cityworks", 1986, "Cityworks provides a platform that helps manage, track, and analyze your infrastructure assets.", "", "https://www.crunchbase.com/organization/cityworks"},
	Company{149, 0, "CivicConnect", 1999, "Civic Connect develops and distributes cloud-based technology solutions for civic organizations.", "", "https://www.crunchbase.com/organization/civic-connect"},
	Company{150, 0, "COBAN Technologies", 2002, "COBAN manufactures state-of-the-art in-car video, body cameras, mobile computers, and evidence management solutions for law enforcement.", "", "https://www.crunchbase.com/organization/coban-technologies"},
	Company{151, 0, "Comcate", 2000, "Comcate develops customer service software solutions for public agencies.", "", "https://www.crunchbase.com/organization/comcate"},
	Company{152, 0, "CyPhy Works", 2008, "CyPhy works develops unmanned air vehicles for search and rescue missions and bridge inspections.", "", "https://www.crunchbase.com/organization/cyphy-works"},
	Company{153, 0, "Department of Better Technology", 2013, "Department of Better Technology makes great software that helps governments and non-profits better serve their communities.", "", "https://www.crunchbase.com/organization/department-of-better-technology"},
	Company{154, 0, "DigitalTown", 1980, "DigitalTown is building the leading platform for online community management powered by content, community and commerce.", "", "https://www.crunchbase.com/organization/digitaltown"},
	Company{155, 0, "EagleEye Intelligence, LLC", 2015, "EagleEye Intelligence provides incident management and aerial platform solutions for public safety and security", "", "https://www.crunchbase.com/organization/409569725"},
	Company{156, 0, "ElectSolve", 2000, "ElectSolve is a full service provider of technical services, consulting and advanced data management solutions.", "", "https://www.crunchbase.com/organization/electsolve"},
	Company{157, 0, "Enigma", 2011, "Enigma sells operational data management and intelligence, and supports the open data community by making its vast library of data public.", "", "https://www.crunchbase.com/organization/enigma"},
	Company{158, 0, "EvoGov", 1997, "EvoGov is an internet provider offering website hosting, design, and other services to many types of organizations.", "", "https://www.crunchbase.com/organization/evogov"},
	Company{159, 0, "Fast Enterprises", 1997, "Fast Enterprises (FAST) provides software and information technology consulting services for government agencies.", "", "https://www.crunchbase.com/organization/409718785"},
	Company{160, 0, "FiscalNote", 2013, "FiscalNote uses artificial intelligence and big data to deliver predictive analytics of governmental action to determine its impact.", "", "https://www.crunchbase.com/organization/fiscalnote"},
	Company{161, 0, "GovDelivery", 2000, "GovDelivery enables public sector organizations to connect with more people and to get those people to take action.", "", "https://www.crunchbase.com/organization/govdelivery"},
	Company{162, 0, "Junar", 2010, "Junar offers a cloud-based open data platform allowing businesses to free their data to drive opportunities, collaboration and transparency.", "", "https://www.crunchbase.com/organization/junar"},
	Company{163, 0, "Meter Feeder", 2014, "Meter Feeder provides a low cost payment and enforcement solution for small to mid-sized governments.", "", "https://www.crunchbase.com/organization/meter-feeder"},
	Company{164, 0, "MetroTech", 2011, "MetroTech Net is a software-analytics company, focused on the integration of new technologies for the smarter city.", "", "https://www.crunchbase.com/organization/metrotech"},
	Company{165, 0, "mySidewalk", 2010, "mySidewalk is a city intelligence tool that helps local analysts track, analyze, and communicate progress on citywide goals.", "", "https://www.crunchbase.com/organization/mysidewalk"},
	Company{166, 0, "PingThings", 2014, "PingThings is a venture-backed software company, brings best-of-breed big-data software and data science techniques to the utility industry.", "", "https://www.crunchbase.com/organization/pingthings"},
	Company{167, 0, "QScend Technologies, Inc.", 1998, "QScend Technologies provides Web-based software solutions and services for municipalities.", "", "https://www.crunchbase.com/organization/qscend-technologies-inc"},
	Company{168, 0, "Recovers", 2011, "Recovers is a software company making disaster preparedness and recovery smarter.", "", "https://www.crunchbase.com/organization/recovers"},
	Company{169, 0, "Seneca Systems", 2014, "Seneca Systems provides next-generation software purpose-built for government.", "", "https://www.crunchbase.com/organization/seneca-systems"},
	Company{170, 0, "Twenty First Century Utilities", 2015, "Twenty First Century Utilities, LLC is working to transform regulated utilities with a twenty first century model.", "", "https://www.crunchbase.com/organization/twenty-first-century-utilities"},
	Company{171, 0, "WiredBlue", 2011, "WiredBlue is an advance technology company focused specifically on law enforcement.", "", "https://www.crunchbase.com/organization/wiredblue"},
	Company{172, 0, "2FA", 2006, "2FA Inc. is a veteran-owned, cybersecurity company created on the single vision of simplifying authentication.", "", "https://www.crunchbase.com/organization/2fa"},
	Company{173, 0, "AutoGrid", 2011, "AutoGrid Systems organizes energy data and employs big data analytics to generate real-time predictions that create actionable data.", "", "https://www.crunchbase.com/organization/autogrid"},
	Company{174, 0, "BasicGov Systems", 2006, "Empowering governments to efficiently serve. BasicGov is a cloud-based SaaS solution making government operations more efficient.", "", "https://www.crunchbase.com/organization/basicgov"},
	Company{175, 0, "BS&A Software", 1987, "BS&A Software provides integrated systems of financial management and property tax products designed specifically for municipalities.", "", "https://www.crunchbase.com/organization/bsa-software"},

	Company{177, 0, "Captricity", 2011, "Captricity is a cloud-based service that converts information on paper to digital data rapidly and effectively.", "", "https://www.crunchbase.com/organization/captricity"},
	Company{178, 0, "CityScan", 2011, "CityScan provides street-level intelligence services for municipalities in the United States.", "", "https://www.crunchbase.com/organization/cityscan"},
	Company{179, 0, "Civil Maps", 2015, "Civil Maps provides a sensor-agnostic cognition platform that enables real-time mapping+localization and crowdsourcing.", "", "https://www.crunchbase.com/organization/civilmaps"},
	Company{180, 0, "Civinomics", 2011, "Civinomics is a community-based web platform that provides crowdsourced solutions to local problems and fostering engagement.", "", "https://www.crunchbase.com/organization/civinomics"},
	Company{181, 0, "Connected Bits", 0, "ConnectedBits develops mobile applications to connect governments and other organizations with their communities.", "", "https://www.crunchbase.com/organization/connectedbits"},
	Company{182, 0, "CrimeStar", 1999, "Crimestar is an information management tools for law enforcement and public safety.", "", "https://www.crunchbase.com/organization/crimestar"},
	Company{183, 0, "DataMade", 2012, "DataMade is a civic technology company that builds open source technology using open data.", "", "https://www.crunchbase.com/organization/datamade"},
	Company{184, 0, "DoubleMap", 2009, "DoubleMap is the provider of an innovative Automatic Vehicle Location (AVL) platform for university and public transit systems.", "", "https://www.crunchbase.com/organization/doublemap"},
	Company{185, 0, "Dropcountr", 0, "Dropcountr connects people and their utilities on the mobile devices they use everyday.", "", "https://www.crunchbase.com/organization/dropcountr"},
	Company{186, 0, "eGov Strategies", 2001, "eGov Strategies provides governments with enterprise payment services, content management and additional interactive service delivery tools.", "", "https://www.crunchbase.com/organization/egov-strategies"},
	Company{187, 0, "EngagePoint", 2007, "EngagePoint is an IT service company that accelerates the government’s transition to a modern, integrated, sustainable enterprise.", "", "https://www.crunchbase.com/organization/engagepoint"},
	Company{188, 0, "ePermitHub", 0, "ePermitHub is a web-based portal based in Florida that allows its users to track, review, and inspect their permits.", "", "https://www.crunchbase.com/organization/epermithub"},
	Company{189, 0, "FireStop", 2013, "FireStop is a cloud-based response software that helps firefighters share critical response information and leverage data in real time.", "", "https://www.crunchbase.com/organization/firestop"},
	Company{190, 0, "GovInvest", 2014, "GovInvest helps governments visualize and understand complex actuarial data.", "", "https://www.crunchbase.com/organization/govinvest"},
	Company{191, 0, "iWorQ Systems", 2001, "iWorQ Systems provides municipal management software.", "", "https://www.crunchbase.com/organization/iworq"},
	Company{192, 0, "LocalData", 0, "LocalData makes digital tools to collect and analyze information about urban infrastructure.", "", "https://www.crunchbase.com/organization/localdata"},
	Company{193, 0, "Localisto", 2011, "Localisto is a new way to discover your community. Find topics and events that matter to you with location-based search and preference", "", "https://www.crunchbase.com/organization/localisto"},
	Company{194, 0, "Loci Controls", 2013, "Loci Controls is a developer of wireless sensor and actor network devices for optimizing the extraction of methane from landfills.", "", "https://www.crunchbase.com/organization/loci-controls"},
	Company{195, 0, "Loveland Technologies", 2012, "LOVELAND Technologies develops crowdfunding and social mapping systems, with a unique creative culture and brand that mixes virtualand real.", "", "https://www.crunchbase.com/organization/loveland"},
	Company{196, 0, "Measured Voice", 2010, "Measured Voice is a social media management tool for government.", "", "https://www.crunchbase.com/organization/measuredvoice"},
	Company{197, 0, "Metropia", 2012, "Metropia optimizes transportation systems by using behavioral economics to influence drivers to shift departure times, routes and modes.", "", "https://www.crunchbase.com/organization/metropia"},

	Company{199, 0, "Municibid", 2006, "Municibid is a compliant, online auction and marketing platform.", "", "https://www.crunchbase.com/organization/municibid"},
	Company{200, 0, "MuniLogic Municipal Management Software", 1973, "MuniLogic provides property management and administration software.", "", "https://www.crunchbase.com/organization/munilogic"},
	Company{201, 0, "MuniRent", 2014, "MuniRent helps government agencies share heavy duty equipment internally within their fleet.", "", "https://www.crunchbase.com/organization/munirent"},
	Company{202, 0, "Opower", 2007, "Opower is a SaaS-based customer engagement and energy efficiency company providing the tools consumers need to make better energy decisions.", "", "https://www.crunchbase.com/organization/opower"},

	Company{204, 0, "OppSites", 2014, "OppSites drives economic development by connecting Cities and Real Estate Professionals in a single marketplace", "", "https://www.crunchbase.com/organization/oppsites"},
	Company{205, 0, "Peak Democracy", 2007, "Peak Democracy offers a cloud-based online civic engagement platform called Open Town Hall.", "", "https://www.crunchbase.com/organization/peak-democracy"},
	Company{206, 0, "Placemeter", 2012, "Placemeter is an urban intelligence platform. We quantify modern cities worldwide. The answers you need are all around you.", "", "https://www.crunchbase.com/organization/placemeter"},
	Company{207, 0, "ProductBio", 2012, "ProductBio is a database of product category information for the procurement process.", "", "https://www.crunchbase.com/organization/productbio"},
	Company{208, 0, "Revelstone Labs", 2010, "Revelstone data analytics and reporting platform scaled for small and medium sized jurisidictions.", "", "https://www.crunchbase.com/organization/revelstone"},
	Company{209, 0, "SABR", 2015, "SABR monitors multiple blockchains to identify and locate criminal activity.", "", "https://www.crunchbase.com/organization/sabrio"},
	Company{210, 0, "SnapSense", 2011, "To help community leaders lead by harnessing big data at a local level.", "", "https://www.crunchbase.com/organization/snapsense"},
	Company{211, 0, "StreetCred Software", 2012, "Fugitive and offender case management", "", "https://www.crunchbase.com/organization/streetcred-software"},
	Company{212, 0, "TransparaGov", 2011, "TransparaGov provides analytical, management, and outcomes measurement software to governments.", "", "https://www.crunchbase.com/organization/transparagov"},
	Company{213, 0, "Urban Engines", 2014, "Urban Engines combines big data and spatial analytics to improve urban mobility.", "", "https://www.crunchbase.com/organization/urban-engines"},
}

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve fronted static files
	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		// Our API will consit of just two routes
		// /companies - which will retrieve a list of companies a user can see
		// /companies/like/:companyID - which will capture likes sent to a particular company
		api.GET("/companies", CompanyHandler)
		// api.POST("/companies/like/:companyID", LikeCompany)
	}

	// Start and run the server
	router.Run(":3000")
}

// CompanyHandler retrieves a list of available companies
func CompanyHandler(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, companies)
}
