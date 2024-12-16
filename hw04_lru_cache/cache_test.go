package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)

		c.Clear()
		val, ok = c.Get("aaa")
		require.False(t, ok)
		require.Equal(t, nil, val)
	})

	t.Run("purge logic", func(t *testing.T) {
		c := NewCache(3)
		wasInCache := c.Set("first", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("second", 2)
		require.False(t, wasInCache)

		wasInCache = c.Set("third", 3)
		require.False(t, wasInCache)

		wasInCache = c.Set("fourth", 4)
		require.False(t, wasInCache)

		// перед добавлением третьего элемента вытолкнетcя первый (key=first), который был добавлен,
		// на его место встанет первый (key=third)
		val, ok := c.Get("first")
		require.False(t, ok)
		require.Nil(t, val)

		// выталкиваем самый тухлый
		c = NewCache(3)
		// используем какие-то ключи несколько раз, чтобы появился один тухлый ключ
		wasInCache = c.Set("first", 1)
		require.False(t, wasInCache)

		wasInCache = c.Set("second", 22)
		require.False(t, wasInCache)

		wasInCache = c.Set("third", 33)
		require.False(t, wasInCache)

		wasInCache = c.Set("second", 222)
		require.True(t, wasInCache)

		wasInCache = c.Set("third", 3)
		require.True(t, wasInCache)

		wasInCache = c.Set("second", 2)
		require.True(t, wasInCache)

		// за время добавления новых элементов должен был протухнуть элемент с key=first
		// добавляем новый
		wasInCache = c.Set("fresh", 7)
		require.False(t, wasInCache)

		// ищем с key=first, который должен был пропасть
		val, ok = c.Get("first")
		require.False(t, ok)
		require.Nil(t, val)
	})
}

func TestCacheMultithreading(_ *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
