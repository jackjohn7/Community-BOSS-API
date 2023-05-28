<script lang="ts">

    type tabItem = {
        name: string,
        path: string,
    };

    type tab = {
        name: string,
        collapsed: boolean,
        tabItems: tabItem[],
    };

    const tabs: tab[] = [
        {
            name: "Start Here",
            collapsed: false,
            tabItems: [
                {
                    name: "Getting Started",
                    path: "/[lang]/[version]/getting_started"
                },
                {
                    name: "Why Make This?",
                    path: "/[lang]/[version]/why"
                },
            ],
        },
    ];

    function getHref(path: string): string {
        return path.replace("[lang]", "en").replace("[version]", "v1");
    }

</script>

<div id="toc-pane" class="flex flex-col bg-blue-900 border-r-4 border-red-800 w-[15%] overflow-scroll">
    {#each tabs as tab}
        <div class="flex flex-col px-2">
            <!-- Add expand/collapse symbol later -->
            <div class="text-white text-xl">{tab.name}</div>
            {#if !tab.collapsed}
                {#each tab.tabItems as item}
                    <a class="text-red-400 px-2" href={getHref(item.path)}>{item.name}</a>
                {/each}
            {/if}
        </div>
    {/each}
</div>