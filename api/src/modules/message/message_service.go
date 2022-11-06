package message

import (
	"context"
	"fmt"
	"log"
	"simple_sns_api/ent"
	"simple_sns_api/ent/message"
	"simple_sns_api/src/db"
	"strings"

	"github.com/google/uuid"
)

type MessageService struct{}

func (s MessageService) FindLatestMessages(ctx context.Context, roomIds []uuid.UUID) ([]*ent.Message, error) {
	roomIdStrings := make([]string, len(roomIds))
	for i, e := range roomIds {
		roomIdStrings[i] = "\"" + e.String() + "\""
	}
	sql := `
		SELECT MAX(id)
		FROM messages
		WHERE room_messages IN (%s)
		GROUP BY room_messages
	`
	rows, err := db.Client.QueryContext(ctx, fmt.Sprintf(sql, strings.Join(roomIdStrings, ",")))
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	latestMessageIds := make([]int, len(roomIds))
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		latestMessageIds = append(latestMessageIds, id)
	}

	if err != nil {
		return nil, err
	}

	return db.Client.Message.Query().
		Where(message.IDIn(latestMessageIds...)).
		Order(ent.Desc(message.FieldCreatedAt)).
		All(ctx)
}
