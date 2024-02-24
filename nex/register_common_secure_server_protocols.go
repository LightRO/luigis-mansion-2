package nex

import (
	"github.com/PretendoNetwork/luigis-mansion-2/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/matchmake-extension"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/matchmake-extension"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/match-making-ext"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/match-making-ext"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/nat-traversal"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/secure-connection"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	common_secure.NewCommonProtocol(secureProtocol)

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	common_match_making.NewCommonProtocol(matchMakingProtocol)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.GetUserFriendPIDs(globals.GetUserFriendPIDs)
}
