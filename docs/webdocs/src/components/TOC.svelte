<script lang="ts">

    import {tabs} from "../data/data";

    function getHref(path: string): string {
        return path.replace("[lang]", "en").replace("[version]", "v1");
    }

</script>

<div id="toc-pane" class="flex flex-col bg-[#1a252f] border-r-4 border-red-800 w-[15%] min-w-[15%] max-w-[15%] overflow-scroll py-2">
    {#each tabs as tab}
        <div class="flex flex-col px-2">
            <!-- Add expand/collapse symbol later -->
            <div class="text-white text-xl flex flex-row justify-between">
                {tab.name}
                <button on:click={() => {
                    console.log("test")
                    tab.collapsed = !tab.collapsed
                    }} class="rounded-md bg-gray-500 px-2">{tab.collapsed ? ">" : "v"}</button>
            </div>
            {#if !tab.collapsed}
                {#each tab.tabItems as item}
                    <a class="text-red-400 px-2" href={getHref(item.path)}>{item.name}</a>
                {/each}
            {/if}
        </div>
    {/each}
</div>
