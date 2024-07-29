package template

import (
	"github.com/diamondburned/arikawa/v3/discord"
	statestore "github.com/diamondburned/arikawa/v3/state/store"
)

const MaxKVValueLength = 16 * 1024
const MaxKVKeyLength = 256

type ContextProvider interface {
	ProvideFuncs(funcs map[string]interface{})
	ProvideData(data map[string]interface{})
}

type InteractionProvider struct {
	state       *statestore.Cabinet
	interaction *discord.InteractionEvent
}

func NewInteractionProvider(state *statestore.Cabinet, interaction *discord.InteractionEvent) *InteractionProvider {
	return &InteractionProvider{
		state:       state,
		interaction: interaction,
	}
}

func (p *InteractionProvider) ProvideFuncs(funcs map[string]interface{}) {}

func (p *InteractionProvider) ProvideData(data map[string]interface{}) {
	data["Interaction"] = NewInteractionData(p.state, p.interaction)

	guildData := NewGuildData(p.state, p.interaction.GuildID, nil)
	data["Guild"] = guildData
	data["Server"] = guildData

	data["Channel"] = NewChannelData(p.state, p.interaction.ChannelID, nil)
}

type GuildProvider struct {
	state   *statestore.Cabinet
	guildID discord.GuildID
	guild   *discord.Guild
}

func NewGuildProvider(state *statestore.Cabinet, guildID discord.GuildID, guild *discord.Guild) *GuildProvider {
	return &GuildProvider{
		state:   state,
		guildID: guildID,
		guild:   guild,
	}
}

func (p *GuildProvider) ProvideFuncs(funcs map[string]interface{}) {}

func (p *GuildProvider) ProvideData(data map[string]interface{}) {
	guildData := NewGuildData(p.state, p.guildID, p.guild)
	data["Guild"] = guildData
	data["Server"] = guildData
}

type ChannelProvider struct {
	state     *statestore.Cabinet
	channelID discord.ChannelID
	channel   *discord.Channel
}

func NewChannelProvider(state *statestore.Cabinet, channelID discord.ChannelID, channel *discord.Channel) *ChannelProvider {
	return &ChannelProvider{
		state:     state,
		channelID: channelID,
		channel:   channel,
	}
}

func (p *ChannelProvider) ProvideFuncs(funcs map[string]interface{}) {}

func (p *ChannelProvider) ProvideData(data map[string]interface{}) {
	data["Channel"] = NewChannelData(p.state, p.channelID, p.channel)
}

/* type KVProvider struct {
	guildID      string
	kvStore      store.KVEntryStore
	maxGuildKeys int
}

func NewKVProvider(guildID string, kvStore store.KVEntryStore, maxGuildKeys int) *KVProvider {
	return &KVProvider{
		guildID:      guildID,
		kvStore:      kvStore,
		maxGuildKeys: maxGuildKeys,
	}
}

func (p *KVProvider) ProvideFuncs(funcs map[string]interface{}) {
	funcs["kvSet"] = p.setKey
	funcs["kvGet"] = p.getKey
	funcs["kvIncrease"] = p.increaseKey
	funcs["kvDelete"] = p.deleteKey
	funcs["kvSearch"] = p.searchKeys
}

func (p *KVProvider) ProvideData(data map[string]interface{}) {}

func (kv *KVProvider) getKey(key string) (string, error) {
	entry, err := kv.kvStore.GetKVEntry(context.TODO(), kv.guildID, key)
	if err != nil {
		if err == store.ErrNotFound {
			return "", nil
		}
		return "", err
	}
	return entry.Value, nil
}

func (kv *KVProvider) setKey(key string, value string) error {
	if len(key) > MaxKVKeyLength {
		return fmt.Errorf("key exceeds maximum length of %d", MaxKVKeyLength)
	}
	if len(value) > MaxKVValueLength {
		return fmt.Errorf("value exceeds maximum length of %d", MaxKVValueLength)
	}

	if err := kv.checkKeyCountLimit(); err != nil {
		return err
	}

	err := kv.kvStore.SetKVEntry(context.TODO(), model.KVEntry{
		GuildID:   kv.guildID,
		Key:       key,
		Value:     value,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (kv *KVProvider) increaseKey(key string, delta int) (string, error) {
	if len(key) > MaxKVKeyLength {
		return "", fmt.Errorf("key exceeds maximum length of %d", MaxKVKeyLength)
	}

	if err := kv.checkKeyCountLimit(); err != nil {
		return "", err
	}

	entry, err := kv.kvStore.IncreaseKVEntry(context.TODO(), model.KVEntryIncreaseParams{
		GuildID:   kv.guildID,
		Key:       key,
		Delta:     delta,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return entry.Value, nil
}

func (kv *KVProvider) deleteKey(key string) (string, error) {
	entry, err := kv.kvStore.DeleteKVEntry(context.TODO(), kv.guildID, key)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return entry.Value, nil
}

func (kv *KVProvider) searchKeys(pattern string) (map[string]string, error) {
	entries, err := kv.kvStore.SearchKVEntries(context.TODO(), kv.guildID, pattern)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(entries))
	for _, entry := range entries {
		result[entry.Key] = entry.Value
	}

	return result, nil
}

func (kv *KVProvider) checkKeyCountLimit() error {
	entryCount, err := kv.kvStore.CountKVEntries(context.TODO(), kv.guildID)
	if err != nil {
		return fmt.Errorf("failed to count KV keys: %w", err)
	}

	if int(entryCount) >= kv.maxGuildKeys {
		return fmt.Errorf("maximum number of keys reached: %d", kv.maxGuildKeys)
	}

	return nil
} */
