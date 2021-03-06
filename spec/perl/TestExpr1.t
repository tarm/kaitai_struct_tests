package spec::perl::TestExpr1;

use strict;
use warnings;
use base qw(Test::Class);
use Test::More;
use Expr1;

sub test_expr_1: Test(4) {
    my $r = Expr1->from_file('src/str_encodings.bin');

    is($r->len_of_1(), 10, 'Equals');
    is($r->len_of_1_mod(), 8, 'Equals');
    is($r->str1(), 'Some ASC', 'Equals');
    is($r->str1_len(), 8, 'Equals');
}

Test::Class->runtests;
