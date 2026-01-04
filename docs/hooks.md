# Lifecycle Hooks

Go State Flow automatically discovers methods on your struct that correspond to state changes. This is where you implement your business logic.

## Hook Discovery

The library uses reflection (cached for O(1) runtime) to find methods named after your states.

**Naming Convention:**
State names are normalized to CamelCase.
- `paid` -> `Paid`
- `in_progress` -> `InProgress`

### 1. Guard Hook (`Can<State>`)

Executed **before** a transition occurs.
- **Purpose**: Validation, permission checks.
- **Signature**: `func() error`
- **Behavior**: If it returns an error, the transition is **aborted**.

```go
func (o *Order) CanPaid() error {
    if o.Amount == 0 {
        return fmt.Errorf("cannot pay empty order")
    }
    return nil
}
```

### 2. Exit Hook (`OnExit<State>`)

Executed when **leaving** a state.
- **Purpose**: Cleanup, releasing locks, stopping timers.
- **Signature**: `func()`

```go
func (o *Order) OnExitDraft() {
    fmt.Println("Leaving draft mode. Order is now immutable.")
}
```

### 3. Entry Hook (`OnEnter<State>`)

Executed when **entering** a state (after the state field has been updated).
- **Purpose**: Triggering side effects, sending notifications, updating other systems.
- **Signature**: `func()`

```go
func (o *Order) OnEnterPaid() {
    emailService.SendReceipt(o.ID)
}
```

## Execution Order

When you call `machine.Transition("target")`:

1.  Check if `current -> target` is derived from struct tags.
2.  Call `CanTarget()` (if exists). If error, STOP.
3.  Call `OnExitCurrent()` (if exists).
4.  Update struct field value to `"target"`.
5.  Call `OnEnterTarget()` (if exists).
