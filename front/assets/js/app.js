var testApp = angular.module('testApp', ['ngRoute']);

testApp.config(function($routeProvider) {
    $routeProvider
        .when('/', {
            templateUrl: 'assets/html/test.html',
            controller: 'testController'
        })
        .when('/get', {
            templateUrl: 'assets/html/get.html',
            controller: 'getController'
        })
        .when('/post', {
            templateUrl: 'assets/html/post.html',
            controller: 'postController'
        });
});

testApp.controller('testController', function($scope) {
    $scope.pageClass = 'test';
});