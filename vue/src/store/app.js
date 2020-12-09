module.exports = {
  types: [
    // this line is used by starport scaffolding
		{ type: "comment", fields: ["body", "postID", ] },
    { type: "post", fields: ["title", "body", ] },
    { type: "contact", fields: ["pseudo"] },
    { type: "like/post", fields: ["like","postID"] },
    { type: "like/comment", fields: ["like","commentID"] }
  ],
};
