// Code generated by ent, DO NOT EDIT.

package tweet

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mayocream/twitter/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldID, id))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldContent, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUpdatedAt, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Tweet {
	return predicate.Tweet(sql.FieldContainsFold(FieldContent, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Tweet {
	return predicate.Tweet(sql.FieldLTE(FieldUpdatedAt, v))
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newAuthorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLikedBy applies the HasEdge predicate on the "liked_by" edge.
func HasLikedBy() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, LikedByTable, LikedByPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLikedByWith applies the HasEdge predicate on the "liked_by" edge with a given conditions (other predicates).
func HasLikedByWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newLikedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRetweetedBy applies the HasEdge predicate on the "retweeted_by" edge.
func HasRetweetedBy() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, RetweetedByTable, RetweetedByPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRetweetedByWith applies the HasEdge predicate on the "retweeted_by" edge with a given conditions (other predicates).
func HasRetweetedByWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newRetweetedByStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParentTweet applies the HasEdge predicate on the "parent_tweet" edge.
func HasParentTweet() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTweetTable, ParentTweetColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentTweetWith applies the HasEdge predicate on the "parent_tweet" edge with a given conditions (other predicates).
func HasParentTweetWith(preds ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newParentTweetStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReplies applies the HasEdge predicate on the "replies" edge.
func HasReplies() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, RepliesTable, RepliesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRepliesWith applies the HasEdge predicate on the "replies" edge with a given conditions (other predicates).
func HasRepliesWith(preds ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newRepliesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMentions applies the HasEdge predicate on the "mentions" edge.
func HasMentions() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, MentionsTable, MentionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMentionsWith applies the HasEdge predicate on the "mentions" edge with a given conditions (other predicates).
func HasMentionsWith(preds ...predicate.User) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newMentionsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasHashtags applies the HasEdge predicate on the "hashtags" edge.
func HasHashtags() predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, HashtagsTable, HashtagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasHashtagsWith applies the HasEdge predicate on the "hashtags" edge with a given conditions (other predicates).
func HasHashtagsWith(preds ...predicate.Hashtag) predicate.Tweet {
	return predicate.Tweet(func(s *sql.Selector) {
		step := newHashtagsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Tweet) predicate.Tweet {
	return predicate.Tweet(sql.NotPredicates(p))
}
