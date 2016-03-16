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

testApp.controller('getController', function($scope, $http) {
    $scope.pageClass = 'get';

        $scope.getdata = function() {
            $scope.data = [];
            $http.get('api/v1/user/' + $scope.input_name)
                .success(function(result) {
                    $scope.datas = result
                    assert(result)
                }
            );
        };
});

testApp.controller('postController', function($scope, $http) {
    $scope.pageClass = 'post';
});