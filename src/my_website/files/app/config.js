angular
	.module('rudaloo')	
	.config(['$routeProvider', function($routeProvider) {
		$routeProvider
		.when('/', {
			templateUrl: '/app/views/index_content.html'
		})
		.when("/about", {
			templateUrl: '/app/views/about_contents.html'
		})
		.when("/contact", {
			templateUrl: '/app/views/contact_contents.html'
		})
		.when("/projects", {
			templateUrl: '/app/views/projects_contents.html'
		})
		.when("/blog", {
			templateUrl: '/app/views/blog_contents.html',
		})
		.when("/chooseblog/:blogname", {
			templateUrl: function(params) {
				return '/app/views/blogs/' + params.blogname;
			},
		})
		.otherwise({redirectTo: '/'});
	}])