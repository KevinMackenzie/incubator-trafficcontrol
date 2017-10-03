package Extensions::TestExtension::Foo;

#package Extensions::TestExtension

#TODO add license
#
#


#sub new {
#	my $self = {};
#	my $class = shift;
#	return ( bless( $self, $class ) );
#}
#
#sub define {
#	my $self = shift;
#	my $r = shift;
#
#	$namespace = "Kevin";
#	$r->get('/unique_name')->over( authenticated => 1, not_ldap => 1)->to( 'foo#hello', namespace => $namespace );
#}
use strict;
use warnings;

use Mojo::Base 'Mojolicious::Controller';

sub hello {
        my $self = shift;

        open( my $fh, '>>', '/tmp/report.txt');
        print $fh "Call to hello\n";
        close $fh;

        $self->render(text => 'Hello Extension!');
}

1
