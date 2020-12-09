package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers blog-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	// this line is used by starport scaffolding
	r.HandleFunc("/blog/comment", listCommentHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/comment", createCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blog/post", listPostHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/post", createPostHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blog/like/comment", listLikeCommentHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/like/comment", createLikeCommentHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blog/like/post", listLikePostHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/like/post", createLikePostHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/blog/contact", listContactHandler(cliCtx, "blog")).Methods("GET")
	r.HandleFunc("/blog/contact", createContactHandler(cliCtx)).Methods("POST")
}
