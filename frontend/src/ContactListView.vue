<script setup lang="ts">
import { onMounted, ref, watch } from 'vue';
import { getCoreRowModel, getFilteredRowModel, useVueTable } from '@tanstack/vue-table';

interface Contact {
    id: number;
    name: string;
    email: string;
    phone_number: string;
    user_id: number
}

const contacts = ref<Contact[]>([])
const globalFilter = ref('')

// dummy token
const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NjkyMzkxMjcsImlzcyI6IndoYXQtdGhlLWZhY2siLCJ1c2VyX2lkIjoyMH0.oTa2V0oQzc0CXVdzaWhevU85sA9eyWWJlWhCqkXmgMw"


watch(globalFilter, value => {
    table.setGlobalFilter(value)
})

async function getContacts() {
    const res = await fetch("/api/contacts", {
        method: "GET",
        headers: {
            "Authorization": "Bearer " + token
        }
    });
    const data = await res.json();

    contacts.value = data as Contact[]
}

async function deleteContact(contactId: number) {
    await fetch("/api/contacts/" + contactId, {
        method: "DELETE",
        headers: {
            "Authorization": "Bearer " + token
        }
    })

    await getContacts();
}

const table = useVueTable({
    data: contacts,
    columns: [
        {
            header: "Name",
            accessorKey: "name"
        },

        {
            header: "Email",
            accessorKey: "email"
        },

        {
            header: "Phone Number",
            accessorKey: "phone_number"
        },

        {
            header: "User",
            accessorKey: "user_id"
        },

        {
            id: "actions",
            header: "",
            cell: () => null
        }
    ],

    getCoreRowModel: getCoreRowModel(),
    getFilteredRowModel: getFilteredRowModel()
})

onMounted(() => {
    getContacts()
})
</script>

<template>
    <div class="container mx-auto">
        <div class="mt-5">
            <h3 class="text-xl">Contacts List</h3>
        </div>
        <br>
        <div class="w-full">
            <input type="text" class="input input-sm w-sm" placeholder="Search" v-model="globalFilter">
        </div>
        <div class="mt-5 border rounded-box border-base-content/5">
            <table class="table">
                <thead>
                    <tr v-for="headerGroup in table.getHeaderGroups()" :key="headerGroup.id">
                        <th v-for="header in headerGroup.headers">
                            {{ header.column.columnDef.header ?? '' }}
                        </th>
                    </tr>
                </thead>
                <tbody>

                    <tr v-for="row in table.getRowModel().rows" :key="row.id">
                        <td v-for="cell in row.getVisibleCells()" :key="cell.id">
                            <template v-if="cell.column.id == 'actions'">
                                <div class="flex gap-3">
                                    <button class="btn btn-xs btn-error"
                                        @click="deleteContact(cell.row.original.id)">Delete</button>
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