package router

import (
	"github.com/gorilla/mux"
	"github.com/muliCohenGini/go-api/internal/entities/benefit"
	"github.com/muliCohenGini/go-api/internal/entities/event"
	"github.com/muliCohenGini/go-api/internal/entities/friend"
	"github.com/muliCohenGini/go-api/internal/entities/influencer"
	"github.com/muliCohenGini/go-api/internal/entities/user"
	"github.com/muliCohenGini/go-api/internal/utils"
)

func Routes(router *mux.Router) {
	userService := user.NewUserService(user.NewUserRepository())
	influencerService := influencer.NewInfluencerService(influencer.NewInfluencerRepository())
	benefitService := benefit.NewBenefitService(benefit.NewBenefitRepository())
	eventService := event.NewEventsService(event.NewEventRepository())
	friendService := friend.NewFriendsService(friend.NewFriendRepository())

	user.UserHandler(router, userService)
	influencer.InfluencerHandler(router, influencerService)
	benefit.BenefitHandler(router, benefitService)
	event.EventHandler(router, eventService)
	friend.FriendHandler(router, friendService)
	utils.HealthCheckHandler(router)
}
