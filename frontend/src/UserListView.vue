<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { getCoreRowModel, getFilteredRowModel, useVueTable } from '@tanstack/vue-table';

interface User {
    id: number;
    firstname: string;
    lastname: string;
    email: string;
}

const users = ref<User[]>([])
const globalFilter = ref("");

// dummy token
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjkyMzkxMjcsImlzcyI6IndoYXQtdGhlLWZhY2siLCJ1c2VyX2lkIjoyMH0.oTa2V0oQzc0CXVdzaWhevU85sA9eyWWJlWhCqkXmgMw"

watch(globalFilter, value => {
    table.setGlobalFilter(value)
})

async function getUsers() {
    const res = await fetch("/api/users",
        {
            method: "GET",
            headers: {
                "Authorization": "Bearer " + token
            }
        }
    )
    const data = await res.json()
    users.value = data as User[]
}

async function deleteUser(userId: number) {
    await fetch("/api/users/" + userId, {
        method: "DELETE",
        headers: {
            "Authorization": "Bearer " + token
        }
    })

    await getUsers();
}

const table = useVueTable({
    data: users,

    columns: [
        {
            header: "First Name",
            accessorKey: "firstname"
        },

        {
            header: "Last Name",
            accessorKey: "lastname"
        },

        {
            header: "Email",
            accessorKey: "email"
        },
        {
            id: 'actions',
            header: "",
            cell: () => null
        }
    ],


    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel()

});

onMounted(() => {
    getUsers()
})
</script>

<template>
    <div class="container mx-auto">
        <div class="mt-5">
            <h3 class="text-xl">Users List</h3>
        </div>

        <br>

        <div class="w-full">
            <input type="text" v-model="globalFilter" class="input input-sm w-sm" placeholder="Search..">
        </div>

        <div class="mt-5 border rounded-box border-base-content/5">
            <table class="table">
                <thead>
                    <tr v-for="headerGroups in table.getHeaderGroups()" :key="headerGroups.id">
                        <th v-for="header in headerGroups.headers" :key="header.id">
                            {{ header.column.columnDef.header ?? '' }}
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="row in table.getRowModel().rows" :key="row.id">
                        <td v-for="cell in row.getVisibleCells()" :key="cell.id">
                            <template v-if="cell.column.id == 'actions'">
                                <div class="flex gap-3">
                                    <button class="btn btn-xs btn-info">Edit</button>
                                    <button class="btn btn-xs btn-error"
                                        @click="deleteUser(cell.row.original.id)">Delete</button>
                                </div>
                            </template>
                            <template v-else>
                                {{ cell.getValue() }}
                            </template>
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </div>
</template>
