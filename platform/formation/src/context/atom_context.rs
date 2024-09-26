use platform_core::context::EventEmitter;
use std::any::TypeId;

use super::{
    atom::{Atom, WeakAtom},
    node::{AnyNode, NodeValue},
    selector::Selector,
    selector_context::SelectorContext,
    subscription::Subscription,
    AnyContext, Context,
};

pub struct AtomContext<'a, T> {
    ctx: &'a mut Context,
    weak: WeakAtom<T>,
}

impl<V> AnyContext for AtomContext<'_, V> {
    type Result<T> = T;

    fn new_atom<T: NodeValue>(
        &mut self,
        build_atom: impl FnOnce(&mut AtomContext<'_, T>) -> T,
    ) -> Self::Result<Atom<T>> {
        self.ctx.new_atom(build_atom)
    }

    fn read_atom<T: NodeValue>(&self, atom: &Atom<T>) -> &T {
        self.ctx.read_atom(atom)
    }

    fn update_atom<T: NodeValue, R>(
        &mut self,
        atom: &Atom<T>,
        callback: impl FnOnce(&mut T, &mut AtomContext<'_, T>) -> R,
    ) -> Self::Result<R> {
        todo!()
    }

    fn new_selector<T: NodeValue>(
        &mut self,
        build_selector: impl Fn(&mut SelectorContext<'_, T>) -> T + 'static,
    ) -> Self::Result<Selector<T>> {
        todo!()
    }

    fn read_selector<T: NodeValue>(&mut self, atom: &Selector<T>) -> &T {
        todo!()
    }
}

impl<'a, V: NodeValue> AtomContext<'a, V> {
    pub(super) fn new(ctx: &'a mut Context, weak: WeakAtom<V>) -> Self {
        Self { ctx, weak }
    }

    pub fn weak_atom(&self) -> WeakAtom<V> {
        self.weak.clone()
    }
}

impl<'a, V: NodeValue> AtomContext<'a, V> {
    pub fn emit<E>(&mut self, event: E)
    where
        V: EventEmitter<E>,
        E: 'static,
    {
        self.ctx.push_effect(super::Effect::Event {
            emitter: self.weak.key,
            typ: TypeId::of::<E>(),
            payload: Box::new(event),
        });
    }

    pub fn notify(&mut self) {
        if self.ctx.batcher.pending_notifications.insert(self.weak.key) {
            self.ctx.push_effect(super::Effect::Notify {
                emitter: self.weak.key,
            });
        }
    }

    pub fn subscribe<T, N, P>(
        &mut self,
        node: &N,
        mut on_event: impl FnMut(&mut V, N, &P, &mut AtomContext<'_, V>) + 'static,
    ) -> Subscription
    where
        T: EventEmitter<P> + 'static,
        N: AnyNode<T>,
        P: 'static,
    {
        let this = self.weak_atom();
        self.ctx.subscribe_internal(node, move |n, payload, ctx| {
            if let Some(atom) = this.upgrade() {
                atom.update(ctx, |this, atom_context| {
                    on_event(this, n, payload, atom_context)
                });

                true
            } else {
                false
            }
        })
    }
}
