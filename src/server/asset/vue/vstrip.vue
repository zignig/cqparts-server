<template>
    <sui-segment>
            <sui-menu-item @click="viz" :positive="visible" basic compact icon="bars"></sui-menu-item>
        <sui-menu v-show="visible" pointing>
            <a 
                is='sui-menu-item'
                v-for="section in sections"
                :active="sectionActive(section)"
                :key="section"
                :content="section"
                @click="selectSection(section)"
                >
            </a>
        </sui-menu>
        <sui-item-group v-show="visible" divided >
            <sui-button-group v-show="pageCount>1" attached="top" size="small" v-show="visible" >
                <sui-button @click="prevPage" :disabled="currentPage === 0" icon="caret left"></sui-button>
            <sui-button > {{ currentPage + 1 }}/{{ pageCount }} </sui-button>
            <sui-button @click="nextPage" :disabled="currentPage >= pageCount -1" icon="caret right" floated="right"></sui-button>
            </sui-button-group>
            <sui-item :key="model.name" v-for="model in pages">
                <sui-button class="nogap" basic v-on:click="load(model)">
                    <sui-item-image wrapped v-on:click="load(model)" size="tiny" v-show="model.img" :src="model.img">
                    </sui-item-image>
                </sui-button>
                <sui-item-content>
                    <h4 is="sui-header" :color="isActive(model)">
                        <span><sui-icon size="small" name="delete" v-on:click="remove(model)"/></sui-icon></span>
                        {{ model.name }}
                    </h4>
                    <sui-item-extra>
                        <sui-button-group>
                            <sui-button :secondary="isPinned(model)" @click="pin(model)" size="mini" icon="thumbtack" v-on:click=""></sui-button>
                        </sui-button-group>
                    </sui-item-extra>
                </sui-item-content>
            </sui-item>
            <sui-button-group v-show="pageCount>1" attached="top" size="small" v-show="visible" >
                <sui-button @click="prevPage" :disabled="currentPage === 0" icon="caret left"></sui-button>
                <sui-button > {{ currentPage + 1 }}/{{ pageCount }} </sui-button>
                <sui-button @click="nextPage" :disabled="currentPage >= pageCount -1" icon="caret right" floated="right"></sui-button>
            </sui-button-group>
        </sui-item-group>
    </sui-segment>
</template>

<script>
export default {
    data(){
        return {
            page: { 'All':0,'Pinned':0,'Fix':0},
            visible: true,
            active: 'Pinned',
            sections: ['Pinned','All','Fix']
        }
    },
    props: {
        current: "",
        modelList: {
            type: Array,
        },
        size:{
            type: Number,
            required: false,
            default: 6,
        }
    },
    methods:{
        sectionActive: function(name){
            return this.active === name
        },
        selectSection: function(name){
            if (this.visible){
                this.active = name
            }
        },
        viz: function(){
            this.visible = !this.visible;
        },
        nextPage: function(){
            this.page[this.active]++;
        },
        prevPage: function(){
            this.page[this.active]--;
        },
        isPinned: function(obj){
            if (obj.pinned){
                return true
            }
        },
        isActive: function(obj){
            if (this.currenti != null ){
                if (this.current.name  == obj.name){
                    return "green"
                }
                else {
                    return "black"
                }
            }
        },
        pin: function(obj){
            obj.pinned = !obj.pinned
            EventBus.$emit('pin',obj.name);
        },
        load: function(obj){
            EventBus.$emit('select',obj);
        },
        remove : function(index){
            this.modelList.splice(this.modelList.indexOf(index),1);
            EventBus.$emit('delete item',index);
        },
        pinned : function(model){
            return this.modelList.filter(function(model){
                return model.pinned === true
            })
        },
        currentData: function(){
            let data = []
            if ( this.active  == 'Pinned' ){
                data = this.pinned()
            }
            if ( this.active  == 'All' ){
                data = this.modelList
            }
            if ( this.active  == 'Fix' ){
                data = this.modelList
            }
            return data
        },
        paginatedData: function(){
            const start = this.page[this.active]* this.size,
                end = start + this.size;
            data = this.currentData()
            return data.slice(start, end);
        }
    },
    computed:{
        currentPage(){
            return this.page[this.active]
        },
        pageCount(){
            let l = this.currentData().length,
                s = this.size;
            return Math.ceil(l/s);
        },
        pages() {
            return this.paginatedData();
        }
    }
};

</script>

<style>
.nogap {
    padding-top: 0.0 !important;
    padding-right: 0.0 !important;
    padding-left: 0.0 !important;
    padding-bottom: 0.0 !important;
}
</style>
