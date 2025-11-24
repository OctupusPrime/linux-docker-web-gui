<script lang="ts">
  import router from "page";
  import { onDestroy, onMount, type Component } from "svelte";

  type PromiseComponent = {
    name: string;
    loader: Component | undefined;
    component: Promise<any>;
  };

  let layout = $state<PromiseComponent | null>(null);
  let page = $state<PromiseComponent | null>(null);
  let queryParams = $state({});

  interface UseFnParams {
    layout?: Promise<any>;
    layoutName?: string;
    layoutLoader?: Component;
    page: Promise<any>;
    pageName: string;
    pageLoader?: Component;
  }

  function use(params: UseFnParams) {
    if (params.layout && params.layoutName) {
      if (params.layoutName !== layout?.name) {
        layout = {
          name: params.layoutName,
          loader: params.layoutLoader,
          component: params.layout,
        };
      }
    } else if (layout !== null) {
      layout = null;
    }

    if (params.pageName !== page?.name) {
      page = {
        name: params.pageName,
        loader: params.pageLoader,
        component: params.page,
      };
    }
  }

  router("/", (ctx) => {
    use({
      page: import("./index.svelte"),
      pageName: "index",
    });
    queryParams = ctx.params;
  });

  onMount(() => {
    router.start();
  });

  onDestroy(() => {
    router.stop();
  });
</script>

{#snippet promiseRenderer(
  item: PromiseComponent | null,
  children: PromiseComponent | null
)}
  {#if item}
    {#await item.component}
      <item.loader />
    {:then { default: ItemComponent }}
      <ItemComponent {queryParams}>
        {@render promiseRenderer(children, null)}
      </ItemComponent>
    {/await}
  {:else if children}
    {@render promiseRenderer(children, null)}
  {/if}
{/snippet}

{@render promiseRenderer(layout, page)}
