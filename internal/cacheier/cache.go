package cacheier

import (
	"context"
	"encoding/json"
	"errors"
	"finalAssing/internal/models"
	"fmt"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

type RedisConn struct {
	red *redis.Client
}

//go:generate mockgen -source cache.go -destination cache_mock.go -package cacheier

type RedInterface interface {
	AddJobData(ctx context.Context, jobId uint, jobData *models.Job) error
	FetchJobData(ctx context.Context, jobId uint) (*models.Job, error)

	AddOtp(ctx context.Context, otp int, userEmail string) error
	CheckOTP(ctx context.Context, userEmail string, otp int) error
	DeleteOtp(ctx context.Context, userEmail string) error
}

func NewRedConn(client *redis.Client) (RedInterface, error) {
	if client == nil {
		return nil, errors.New("redis client not provided")
	}
	return &RedisConn{red: client}, nil
}

func (r *RedisConn) AddJobData(ctx context.Context, jobId uint, jobData *models.Job) error {
	strId := strconv.FormatUint(uint64(jobId), 10)
	mData, err := json.Marshal(jobData)
	if err != nil {
		log.Error().Err(err).Interface("Job Id", jobId).Msg("unable marshal job Data")

		return err
	}
	err = r.red.Set(ctx, strId, mData, 10*time.Minute).Err()
	if err != nil {
		log.Error().Err(err).Interface("Job Id", jobId).Msg("failure in cache of job")
		return err
	}
	return nil
}

func (r *RedisConn) FetchJobData(ctx context.Context, jobId uint) (*models.Job, error) {
	fmt.Println("job data fetched from Redis")
	strId := strconv.FormatUint(uint64(jobId), 10)
	data, err := r.red.Get(ctx, strId).Result()
	if err != nil {
		log.Error().Err(err).Interface("job Id", jobId).Msg("Can't find job data in redis")
		return nil, err
	}
	jobData := new(models.Job)
	// fmt.Println("DATA::", data)
	err = json.Unmarshal([]byte(data), jobData)
	if err != nil {
		log.Error().Err(err).Interface("Job Id", jobId).Msg("Error in Unmarshaling")
		return nil, err
	}
	return jobData, nil
}

func (r *RedisConn) AddOtp(ctx context.Context, otp int, userEmail string) error {
	err := r.red.Set(ctx, userEmail, otp, 5*time.Minute).Err()
	if err != nil {
		log.Error().Err(err).Interface("user email", userEmail).Msg("failure in cache of OTP")
		return err
	}
	log.Info().Msg("OTP added to redis successfuly!")
	return nil
}

func (r *RedisConn) CheckOTP(ctx context.Context, userEmail string, otp int) error {
	data, err := r.red.Get(ctx, userEmail).Result()
	if err != nil {
		log.Error().Err(err).Msg("Invalid Email")
		return err
	}
	if data != strconv.FormatInt(int64(otp), 10) {
		log.Error().Err(errors.New("OTP didn't match"))
		return err
	}
	log.Info().Msg("OTP Verified")
	return nil
}

func (r *RedisConn) DeleteOtp(ctx context.Context, userEmail string) error {
	_, err := r.red.Del(ctx, userEmail).Result()
	if err != nil {
		log.Error().Err(err).Msg("OTP not deleted from redis")
		return err
	}
	log.Info().Msg("OTP successfuly deleted from redis")
	return nil
}
