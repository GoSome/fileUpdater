<template>
  <!-- Page Wrapper -->
  <div>

    <!-- Sidebar -->
    <ul class="navbar-nav bg-gradient-primary sidebar sidebar-dark accordion" id="accordionSidebar">

      <!-- Sidebar - Brand -->
      <router-link class="sidebar-brand d-flex align-items-center justify-content-center" :to="{name: 'dashboard'}">
        <div class="sidebar-brand-icon rotate-n-15"><i class="fas fa-laugh-wink"></i></div>
        <div class="sidebar-brand-text mx-3">FileUpdater</div>
      </router-link>

      <!-- Divider -->
      <hr class="sidebar-divider my-0">

      <sidebar-item to="dashboard" fa-icon="fa-tachometer-alt" text="控制台"/>
    </ul>
    <!-- End of Sidebar -->

    <div id="content-wrapper" class="d-flex flex-column">
      <div id="content">
        <router-view class="container-fluid"/>
      </div>
      <footer class="sticky-footer bg-white">
        <div class="container my-auto">
          <div class="copyright text-center my-auto">
            <span>Copyright &copy; GoSome</span>
          </div>
        </div>
      </footer>
    </div>
  </div>
</template>

<script>
import SidebarItem from './SidebarItem.vue'


export default {
  components: {
    SidebarItem,
  },
  data() {
    return {
    }
  },
}
</script>
