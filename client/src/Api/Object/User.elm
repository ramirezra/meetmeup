-- Do not manually edit this file, it was auto-generated by dillonkearns/elm-graphql
-- https://github.com/dillonkearns/elm-graphql


module Api.Object.User exposing (..)

import Api.InputObject
import Api.Interface
import Api.Object
import Api.Scalar
import Api.ScalarCodecs
import Api.Union
import Graphql.Internal.Builder.Argument as Argument exposing (Argument)
import Graphql.Internal.Builder.Object as Object
import Graphql.Internal.Encode as Encode exposing (Value)
import Graphql.Operation exposing (RootMutation, RootQuery, RootSubscription)
import Graphql.OptionalArgument exposing (OptionalArgument(..))
import Graphql.SelectionSet exposing (SelectionSet)
import Json.Decode as Decode


{-| -}
id : SelectionSet Api.ScalarCodecs.Id Api.Object.User
id =
    Object.selectionForField "ScalarCodecs.Id" "id" [] (Api.ScalarCodecs.codecs |> Api.Scalar.unwrapCodecs |> .codecId |> .decoder)


{-| -}
username : SelectionSet String Api.Object.User
username =
    Object.selectionForField "String" "username" [] Decode.string


{-| -}
email : SelectionSet String Api.Object.User
email =
    Object.selectionForField "String" "email" [] Decode.string


{-| -}
firstName : SelectionSet String Api.Object.User
firstName =
    Object.selectionForField "String" "firstName" [] Decode.string


{-| -}
lastName : SelectionSet String Api.Object.User
lastName =
    Object.selectionForField "String" "lastName" [] Decode.string


{-| -}
meetups : SelectionSet decodesTo Api.Object.Meetup -> SelectionSet (List decodesTo) Api.Object.User
meetups object_ =
    Object.selectionForCompositeField "meetups" [] object_ (identity >> Decode.list)


{-| -}
createdAt : SelectionSet Api.ScalarCodecs.Time Api.Object.User
createdAt =
    Object.selectionForField "ScalarCodecs.Time" "createdAt" [] (Api.ScalarCodecs.codecs |> Api.Scalar.unwrapCodecs |> .codecTime |> .decoder)


{-| -}
updatedAt : SelectionSet Api.ScalarCodecs.Time Api.Object.User
updatedAt =
    Object.selectionForField "ScalarCodecs.Time" "updatedAt" [] (Api.ScalarCodecs.codecs |> Api.Scalar.unwrapCodecs |> .codecTime |> .decoder)
