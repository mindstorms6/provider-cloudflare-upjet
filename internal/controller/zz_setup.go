// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0


package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	onerequestasset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestasset"
routingcatchall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingcatchall"
securityblocksender "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securityblocksender"
transitconnector "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitconnector"
transforms "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/managed/transforms"
trustaccessidentityprovider "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessidentityprovider"
trustdnslocation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdnslocation"
zonetransfersincoming "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersincoming"
balancermonitor "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancermonitor"
balancerpool "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancerpool"
bucketlifecycle "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketlifecycle"
cache "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/tiered/cache"
lockdown "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/lockdown"
shield "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shield"
image "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/image"
job "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpush/job"
hostnameweb3 "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web3/hostname"
kv "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/kv"
trustdevicepostureintegration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicepostureintegration"
token "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/token"
routingaddress "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingaddress"
bucketeventnotification "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketeventnotification"
download "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/download"
forplatformsdispatchnamespace "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/forplatformsdispatchnamespace"
route "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/route"
trustaccesscustompage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesscustompage"
trustdlppredefinedprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlppredefinedprofile"
scanningexpression "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/content/scanningexpression"
trustaccessservicetoken "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessservicetoken"
trustdevicedefaultprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofile"
trustgatewaylogging "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaylogging"
stream "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/stream"
domain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/pages/domain"
originpulls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/authenticated/originpulls"
securityimpersonationregistry "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securityimpersonationregistry"
credentialcheckrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/leaked/credentialcheckrule"
wangretunnel "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wangretunnel"
limit "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/rate/limit"
tls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/total/tls"
script "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/script"
trustaccessshortlivedcertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessshortlivedcertificate"
shieldoperation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldoperation"
hostname "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/hostname"
wanstaticroute "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wanstaticroute"
trustdevicecustomprofilelocaldomainfallback "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicecustomprofilelocaldomainfallback"
trustdevicedefaultprofilelocaldomainfallback "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofilelocaldomainfallback"
trustdlpentry "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpentry"
trustgatewaycertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaycertificate"
trustgatewayproxyendpoint "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewayproxyendpoint"
onerequest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequest"
domainregistrar "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/registrar/domain"
audiotrack "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/audiotrack"
analyticsrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web/analyticsrule"
trustdevicedefaultprofilecertificates "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofilecertificates"
trustdlpdataset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpdataset"
cachereserve "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/cachereserve"
dnssec "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/dnssec"
zonetransferstsig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransferstsig"
retention "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpull/retention"
application "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/spectrum/application"
crontrigger "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/crontrigger"
trustaccessgroup "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessgroup"
hold "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/hold"
setting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/setting"
tokenapi "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/token"
settingsinternalview "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/settingsinternalview"
balancer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancer"
transitsite "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsite"
wanipsectunnel "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wanipsectunnel"
bucket "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucket"
key "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/key"
agentblockingrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/user/agentblockingrule"
item "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/list/item"
webhook "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/webhook"
trustaccesstag "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesstag"
member "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/member"
shielddiscoveryoperation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shielddiscoveryoperation"
list "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/list"
record "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/record"
roomsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomsettings"
trusttunnelcloudflaredconfig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredconfig"
captionlanguage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/captionlanguage"
subscription "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/subscription"
shieldschema "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldschema"
shieldschemavalidationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldschemavalidationsettings"
healthcheck "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/healthcheck"
manageddomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/manageddomain"
tieredcache "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/regional/tieredcache"
liveinput "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/liveinput"
connectorrules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloud/connectorrules"
settings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/settings"
customdomainworkers "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/customdomain"
trustorganization "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustorganization"
turnapp "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/calls/turnapp"
onerequestmessage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestmessage"
routingrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingrule"
consumer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/queue/consumer"
bucketcors "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketcors"
trustaccessapplication "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessapplication"
trustgatewaysettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaysettings"
tieredcaching "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/argo/tieredcaching"
zonetransfersoutgoing "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersoutgoing"
routingdns "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingdns"
kvnamespace "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/kvnamespace"
trustaccessinfrastructuretarget "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessinfrastructuretarget"
ruleset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/ruleset"
certificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/keyless/certificate"
networkmonitoringrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/networkmonitoringrule"
certificatemtls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/mtls/certificate"
roomrules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomrules"
trustaccesspolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesspolicy"
pack "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/certificate/pack"
database "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/d1/database"
zonetransferspeer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransferspeer"
rulefirewall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/firewall/rule"
variant "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/image/variant"
policy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/notification/policy"
customdomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/customdomain"
hostnameregional "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/regional/hostname"
rule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/access/rule"
ipprefix "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/byo/ipprefix"
transitsiteacl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsiteacl"
scheduledtest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/observatory/scheduledtest"
secret "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/secret"
trustdevicecustomprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicecustomprofile"
trustdevicemanagednetworks "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicemanagednetworks"
cachevariants "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/cachevariants"
credentialcheck "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/leaked/credentialcheck"
transitsitelan "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsitelan"
transitsitewan "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsitewan"
trustdextest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdextest"
trustgatewaypolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaypolicy"
firewall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/firewall"
management "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/bot/management"
filter "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/filter"
project "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/pages/project"
rules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/snippet/rules"
subscriptionzone "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/subscription"
onerequestpriority "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestpriority"
roomevent "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomevent"
trustriskbehavior "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustriskbehavior"
trustriskscoringintegration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustriskscoringintegration"
trusttunnelcloudflaredvirtualnetwork "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredvirtualnetwork"
account "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/account"
policywebhooks "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/notification/policywebhooks"
rulepage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/page/rule"
bucketlock "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketlock"
room "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/room"
deployment "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/deployment"
snippets "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/snippets"
shieldpolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/page/shieldpolicy"
providerconfig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/providerconfig"
queue "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/queue"
config "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/hyperdrive/config"
analyticssite "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web/analyticssite"
trustdlpcustomprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpcustomprofile"
zonetransfersacl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersacl"
securitytrusteddomains "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securitytrusteddomains"
networkmonitoringconfiguration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/networkmonitoringconfiguration"
watermark "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/watermark"
widget "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/turnstile/widget"
normalizationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/url/normalizationsettings"
scriptsubdomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/scriptsubdomain"
trustaccessmtlscertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessmtlscertificate"
routingsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingsettings"
ownershipchallenge "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpush/ownershipchallenge"
_map "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/address/_map"
tlssetting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/hostname/tlssetting"
bucketsippy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketsippy"
zone "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/zone"
user "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/user"
trustaccesskeyconfiguration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesskeyconfiguration"
trustlist "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustlist"
trusttunnelcloudflaredroute "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredroute"
smartrouting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/argo/smartrouting"
hostnamefallbackorigin "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/hostnamefallbackorigin"
trustdeviceposturerule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdeviceposturerule"
shieldoperationschemavalidationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldoperationschemavalidationsettings"
originpullscertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/authenticated/originpullscertificate"
ssl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/ssl"
cacertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/origin/cacertificate"
trustaccessmtlshostnamesettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessmtlshostnamesettings"
trusttunnelcloudflared "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflared"
sfuapp "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/calls/sfuapp"

)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.Setup,
		member.Setup,
		subscription.Setup,
		token.Setup,
		map.Setup,
		shield.Setup,
		shielddiscoveryoperation.Setup,
		shieldoperation.Setup,
		shieldoperationschemavalidationsettings.Setup,
		shieldschema.Setup,
		shieldschemavalidationsettings.Setup,
		tokenapi.Setup,
		smartrouting.Setup,
		tieredcaching.Setup,
		originpulls.Setup,
		originpullscertificate.Setup,
		management.Setup,
		ipprefix.Setup,
		sfuapp.Setup,
		turnapp.Setup,
		pack.Setup,
		connectorrules.Setup,
		account.Setup,
		filter.Setup,
		healthcheck.Setup,
		image.Setup,
		list.Setup,
		queue.Setup,
		ruleset.Setup,
		snippets.Setup,
		stream.Setup,
		user.Setup,
		zone.Setup,
		onerequest.Setup,
		onerequestasset.Setup,
		onerequestmessage.Setup,
		onerequestpriority.Setup,
		scanningexpression.Setup,
		hostname.Setup,
		hostnamefallbackorigin.Setup,
		ssl.Setup,
		database.Setup,
		firewall.Setup,
		record.Setup,
		settings.Setup,
		settingsinternalview.Setup,
		zonetransfersacl.Setup,
		zonetransfersincoming.Setup,
		zonetransfersoutgoing.Setup,
		zonetransferspeer.Setup,
		zonetransferstsig.Setup,
		routingaddress.Setup,
		routingcatchall.Setup,
		routingdns.Setup,
		routingrule.Setup,
		routingsettings.Setup,
		securityblocksender.Setup,
		securityimpersonationregistry.Setup,
		securitytrusteddomains.Setup,
		rulefirewall.Setup,
		tlssetting.Setup,
		config.Setup,
		variant.Setup,
		certificate.Setup,
		credentialcheck.Setup,
		credentialcheckrule.Setup,
		item.Setup,
		balancer.Setup,
		balancermonitor.Setup,
		balancerpool.Setup,
		retention.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		networkmonitoringconfiguration.Setup,
		networkmonitoringrule.Setup,
		transitconnector.Setup,
		transitsite.Setup,
		transitsiteacl.Setup,
		transitsitelan.Setup,
		transitsitewan.Setup,
		wangretunnel.Setup,
		wanipsectunnel.Setup,
		wanstaticroute.Setup,
		transforms.Setup,
		certificatemtls.Setup,
		policy.Setup,
		policywebhooks.Setup,
		scheduledtest.Setup,
		cacertificate.Setup,
		rulepage.Setup,
		shieldpolicy.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		consumer.Setup,
		bucket.Setup,
		bucketcors.Setup,
		bucketeventnotification.Setup,
		bucketlifecycle.Setup,
		bucketlock.Setup,
		bucketsippy.Setup,
		customdomain.Setup,
		manageddomain.Setup,
		limit.Setup,
		hostnameregional.Setup,
		tieredcache.Setup,
		domainregistrar.Setup,
		rules.Setup,
		application.Setup,
		audiotrack.Setup,
		captionlanguage.Setup,
		download.Setup,
		key.Setup,
		liveinput.Setup,
		watermark.Setup,
		webhook.Setup,
		cache.Setup,
		tls.Setup,
		widget.Setup,
		normalizationsettings.Setup,
		agentblockingrule.Setup,
		room.Setup,
		roomevent.Setup,
		roomrules.Setup,
		roomsettings.Setup,
		analyticsrule.Setup,
		analyticssite.Setup,
		hostnameweb3.Setup,
		crontrigger.Setup,
		customdomainworkers.Setup,
		deployment.Setup,
		forplatformsdispatchnamespace.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		scriptsubdomain.Setup,
		secret.Setup,
		trustaccessapplication.Setup,
		trustaccesscustompage.Setup,
		trustaccessgroup.Setup,
		trustaccessidentityprovider.Setup,
		trustaccessinfrastructuretarget.Setup,
		trustaccesskeyconfiguration.Setup,
		trustaccessmtlscertificate.Setup,
		trustaccessmtlshostnamesettings.Setup,
		trustaccesspolicy.Setup,
		trustaccessservicetoken.Setup,
		trustaccessshortlivedcertificate.Setup,
		trustaccesstag.Setup,
		trustdevicecustomprofile.Setup,
		trustdevicecustomprofilelocaldomainfallback.Setup,
		trustdevicedefaultprofile.Setup,
		trustdevicedefaultprofilecertificates.Setup,
		trustdevicedefaultprofilelocaldomainfallback.Setup,
		trustdevicemanagednetworks.Setup,
		trustdevicepostureintegration.Setup,
		trustdeviceposturerule.Setup,
		trustdextest.Setup,
		trustdlpcustomprofile.Setup,
		trustdlpdataset.Setup,
		trustdlpentry.Setup,
		trustdlppredefinedprofile.Setup,
		trustdnslocation.Setup,
		trustgatewaycertificate.Setup,
		trustgatewaylogging.Setup,
		trustgatewaypolicy.Setup,
		trustgatewayproxyendpoint.Setup,
		trustgatewaysettings.Setup,
		trustlist.Setup,
		trustorganization.Setup,
		trustriskbehavior.Setup,
		trustriskscoringintegration.Setup,
		trusttunnelcloudflared.Setup,
		trusttunnelcloudflaredconfig.Setup,
		trusttunnelcloudflaredroute.Setup,
		trusttunnelcloudflaredvirtualnetwork.Setup,
		cachereserve.Setup,
		cachevariants.Setup,
		dnssec.Setup,
		hold.Setup,
		lockdown.Setup,
		setting.Setup,
		subscriptionzone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
