package services

import (
	"fmt"
	"strconv"
	"wuju_data/entity"

	"github.com/bwmarrin/snowflake"
	"github.com/rs/zerolog/log"
)

type snowFlakeService struct {
	node *snowflake.Node
}

func (s *snowFlakeService) Setup(c *entity.Conf) error {
	if len(c.Leagues) == 0 {
		return fmt.Errorf("conf League length = 0")
	}

	//
	var total int64
	for _, v := range c.Leagues {
		c, _ := strconv.ParseInt(v.ID, 10, 64)
		total += c
	}

	machineID := total % 1023
	log.Info().Int64("machine_id", machineID).Send()

	var err error
	s.node, err = snowflake.NewNode(machineID)
	if err != nil {
		return err
	}
	return nil
}

// 生成int64 id
func (s *snowFlakeService) NextID() string {
	return fmt.Sprintf("%d", s.node.Generate().Int64())
}
