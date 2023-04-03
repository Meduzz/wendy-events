# wendy-events
Let your wendy modules listen to events.

## Background

Since wendy core is completely disconnected from any type of connections (somewhat ironic), this type of functionality simply cannot be stored with it. Since that would bloat away one of the core qualities we enjoy when embedding the framework.

## How

For now listening to events are added in a somewhat disconnected fashion, but we try to make up for it with the wrapper class.

Basically, there's a new struct called `EventListener` which will do the logic around events. There's a new api level struct called `Event` which serves as the base abstraction for the event. There´s also a new handler type, called `Listener` (`func(*Event)`).

You tie your listeners to topics with `func (e* EventListener) WithListener(topic string, listener Listener)`.

And then since, we now rely on connections, there's a start method that accepts a `*github.com/Meduzz/rpc.RPC`, that will setup your listeners.
